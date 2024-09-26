variable "docker_host" {
  type = string
  default = "unix:///var/run/docker.sock"
}

variable "api_envs" {
  type = map(string)
  default = {}
}