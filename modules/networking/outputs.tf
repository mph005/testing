output "vm_subnet_id" {
  description = "ID of the VM subnet"
  value       = azurerm_subnet.vm_subnet.id
}

output "bastion_host_id" {
  description = "ID of the Bastion Host"
  value       = azurerm_bastion_host.bastion.id
}

output "vm_nsg_id" {
  description = "ID of the VM NSG"
  value       = azurerm_network_security_group.vm_nsg.id
} 