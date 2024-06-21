provider "docker" {
    host = var.docker_host
}

resource "docker_image" "mongo" {
  name = "mongo:latest"
  keep_locally = false
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