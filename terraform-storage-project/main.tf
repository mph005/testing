resource "azurerm_resource_group" "example" {
  name     = var.resource_group_name
  location = var.location
}

module "storage_account" {
  source                    = "./terraform-storage-module"
  storage_account_name      = var.storage_account_name
  resource_group_name       = azurerm_resource_group.example.name
  location                  = var.location
  account_tier              = var.account_tier
  account_replication_type  = var.account_replication_type
  tags                      = var.tags
}