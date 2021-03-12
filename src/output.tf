output "network" {
  value = module.network.name
}

output "gce_static_ip" {
  value = module.gce.static_ip[local.name]
}

