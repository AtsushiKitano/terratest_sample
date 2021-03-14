locals {
  ingress_sample_firewall_allow_rule = { for v in module.network.firewall.ingress-sample.allow : "allow_rule" => v }
  gce_nics                           = { for v in module.gce.gce.network_interface : "0" => v }
}

output "network_name" {
  value = module.network.vpc.name
}

output "subnetwork_name" {
  value = module.network.subnetwork[local.name].name
}

output "subnetwork_region" {
  value = module.network.subnetwork[local.name].region
}

output "subnetwork_cidr" {
  value = module.network.subnetwork[local.name].ip_cidr_range
}

output "vpc_auto_create_subnetworks" {
  value = module.network.vpc.auto_create_subnetworks
}

output "firewall_name" {
  value = module.network.firewall.ingress-sample.name
}

output "firewall_direction" {
  value = module.network.firewall.ingress-sample.direction
}

output "firewall_src_ranges" {
  value = [for v in module.network.firewall.ingress-sample.source_ranges : v]
}

output "firewall_priority" {
  value = module.network.firewall.ingress-sample.priority
}

output "firewall_allow_rules_protocol" {
  value = local.ingress_sample_firewall_allow_rule.allow_rule.protocol
}

output "firewall_allow_rules_ports" {
  value = [for v in local.ingress_sample_firewall_allow_rule.allow_rule.ports : v]
}

output "gce_name" {
  value = module.gce.gce.name
}

output "gce_machine_type" {
  value = module.gce.gce.machine_type
}

output "gce_zone" {
  value = module.gce.gce.zone
}

output "gce_boot_disk_size" {
  value = module.gce.boot_disk.size
}

output "gce_boot_disk_image" {
  value = module.gce.boot_disk.image
}

output "gce_global_ip" {
  value = [for v in local.gce_nics.0.access_config : v.nat_ip]
}
