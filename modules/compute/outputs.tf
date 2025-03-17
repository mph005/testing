output "vm_ids" {
  description = "IDs of the created VMs"
  value       = azurerm_linux_virtual_machine.vm[*].id
}

output "vm_principal_ids" {
  description = "Principal IDs of the VMs"
  value       = azurerm_linux_virtual_machine.vm[*].identity[0].principal_id
}

output "vm_private_ips" {
  description = "Private IP addresses of the VMs"
  value       = azurerm_network_interface.vm_nic[*].private_ip_address
} 