output "storage_account_id" {
  value = azurerm_storage_account.example.id
}

output "primary_blob_endpoint" {
  value = azurerm_storage_account.example.primary_blob_endpoint
}

output "primary_queue_endpoint" {
  value = azurerm_storage_account.example.primary_queue_endpoint
}

output "primary_table_endpoint" {
  value = azurerm_storage_account.example.primary_table_endpoint
}

output "primary_file_endpoint" {
  value = azurerm_storage_account.example.primary_file_endpoint
}