resource "docker_network" "consul_network" {
  name = "consul_network"
}

resource "docker_image" "consul" {
  name = "consul:1.2.3"
}

resource "docker_container" "consul1" {
  name  = "consul1"
  image = docker_image.consul.latest
  command = ["agent", "-server", "-bootstrap-expect=1"]
  hostname = "consul1"
  publish_all_ports = true
  networks_advanced {
    name = docker_network.consul_network.name
  }
}

output "consul1_docker_id" {
  value = docker_container.consul1.id
}

output "consul1_ip" {
  value = docker_container.consul1.network_data[0].ip_address
}
