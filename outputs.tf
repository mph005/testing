output "resource_group_name" {
  description = "Name of the created resource group"
  value       = azurerm_resource_group.rg.name
}

output "vm_private_ips" {
  description = "Private IP addresses of the VMs"
  value       = module.compute.vm_private_ips
}

output "storage_account_name" {
  description = "Name of the storage account"
  value       = module.storage.storage_account_name
}

output "bastion_host_name" {
  description = "Name of the Bastion Host"
  value       = module.networking.bastion_host_id
}

output "vm_ids" {
  description = "IDs of the created VMs"
  value       = module.compute.vm_ids
} 