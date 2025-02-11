resource "azurerm_storage_account" "example" {
  name                     = azurerm_storage_account.example.name
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = azurerm_storage_account.example.account_tier
  account_replication_type = azurerm_storage_account.example.account_replication_type
}

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