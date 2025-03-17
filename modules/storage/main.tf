resource "azurerm_storage_account" "storage" {
  name                     = "st${var.environment}${random_string.storage_suffix.result}"
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = var.storage_account_tier
  account_replication_type = var.storage_account_replication_type
  tags                     = var.tags

  network_rules {
    default_action = "Deny"
    virtual_network_subnet_ids = [var.vm_subnet_id]
  }
}

resource "random_string" "storage_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "azurerm_storage_container" "blob_container" {
  name                  = "data"
  storage_account_name  = azurerm_storage_account.storage.name
  container_access_type = "private"
}

resource "azurerm_role_assignment" "vm_contributor" {
  count                = 2
  scope                = azurerm_storage_account.storage.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = var.vm_principal_ids[count.index]
} 