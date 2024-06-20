provider "docker" {
    host = var.docker_host
}

resource "docker_image" "nginx" {
  name = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.image_id
  name = "tutorial-nginx"
  ports {
    internal = 80
    external = 8080
  }
}

resource "docker_image" "mongo" {
  name = "mongo:latest"
  keep_locally = false
}

resource "docker_volume" "mongo_volume" {
  name = "mongo_volume"
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
}