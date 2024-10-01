provider "docker" {
    host = var.docker_host
}

resource "docker_image" "mongo" {
  name = "mongo:latest"
}

resource "docker_image" "api" {
  name = "api"
  build {
    context = "${path.module}/../"
    tag = [ "api:latest" ]
  }
}

resource "docker_volume" "mongo_volume" {
  name = "mongo_volume"
}

resource "docker_volume" "mongo_volume2" {
  name = "mongo_volume2"
}

resource "docker_volume" "mongo_volume3" {
  name = "mongo_volume3"
}

resource "docker_network" "mongo-cluster" {
  name = "mongo-cluster"
  driver = "bridge"
}

resource "docker_container" "mongo" {
  image = docker_image.mongo.image_id
  name = "mongo-test"
  ports {
    internal = 27017
    external = 27017
  }
  volumes {
    volume_name = docker_volume.mongo_volume.name
    container_path = "/data/db"
    host_path = "${path.cwd}/mongo/data/db"
  }

  healthcheck {
    test = [ "CMD", "echo", "\"try { rs.status() } catch (err) { rs.initiate({_id:'repl1',members:[{_id:0,host:'mongo-test:27017'}, {_id:1,host:'mongo-repl1:27017'},{_id:2,host:'mongo-repl2:27017'}]}) }\"", "|", "mongosh", "--port", "27017", "--quiet" ]
    interval = "5s"
    timeout = "30s"
    start_period = "0s"
    retries = 30
  }

  networks_advanced {
    name = docker_network.mongo-cluster.name
  }

  command = [ "mongod", "--replSet", "repl1" ]
}

resource "docker_container" "mongo-repl1" {
  image = docker_image.mongo.image_id
  name = "mongo-repl1"
  ports {
    internal = 27017
    external = 3001
  }
  volumes {
    volume_name = docker_volume.mongo_volume2.name
    container_path = "/data/db"
    host_path = "${path.cwd}/mongo/data/db1"
  }

  networks_advanced {
    name = docker_network.mongo-cluster.name
  }
  
  command = [ "mongod", "--replSet", "repl1" ]
}

resource "docker_container" "mongo-repl2" {
  image = docker_image.mongo.image_id
  name = "mongo-repl2"
  ports {
    internal = 27017
    external = 3002
  }
  volumes {
    volume_name = docker_volume.mongo_volume3.name
    container_path = "/data/db"
    host_path = "${path.cwd}/mongo/data/db2"
  }

  networks_advanced {
    name = docker_network.mongo-cluster.name
  }
  
  command = [ "mongod", "--replSet", "repl1" ]
}

resource "docker_container" "mongo-init" {
  image = docker_image.mongo.image_id
  name = "mongo-init"
  depends_on = [ docker_container.mongo, docker_container.mongo-repl1, docker_container.mongo-repl2 ]

  networks_advanced {
    name = docker_network.mongo-cluster.name
  }
  
  restart = "on-failure"
  command = [ "mongosh", "--host", "mongo-test:27017", "--eval",  "'try{ rs.status(); } catch(err){ rs.initiate({_id:'repl1',members:[{_id:0,host:'mongo-test:27017'}, {_id:1,host:'mongo-repl1:27017'},{_id:2,host:'mongo-repl2:27017'}]})}; rs.status()'" ]
}

resource "docker_container" "api" {
  name = "api"
  image = docker_image.api.image_id
  depends_on = [ docker_container.mongo, docker_container.mongo-repl1, docker_container.mongo-repl2 ]
  networks_advanced {
    name = docker_network.mongo-cluster.name
  }

  env = toset([for k, v in var.api_envs : "${k}=${v}"])

  ports {
    internal = 8090
    external = 8090
  }

  restart = "always"
}