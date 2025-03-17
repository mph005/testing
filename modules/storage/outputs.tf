output "storage_account_name" {
  description = "Name of the storage account"
  value       = azurerm_storage_account.storage.name
}

output "storage_account_key" {
  description = "Primary access key for the storage account"
  value       = azurerm_storage_account.storage.primary_access_key
  sensitive   = true
}

output "container_name" {
  description = "Name of the blob container"
  value       = azurerm_storage_container.blob_container.name
} 