# Storage module for Azure VM Storage Project
# This module creates the storage infrastructure including storage account and blob container

# Create a storage account with a random suffix for uniqueness
# Network rules are configured to only allow access from the VM subnet
resource "azurerm_storage_account" "storage" {
  name                     = "st${var.environment}${random_string.storage_suffix.result}"
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = var.storage_account_tier
  account_replication_type = var.storage_account_replication_type
  tags                     = var.tags

  # Configure network rules to restrict access
  network_rules {
    default_action             = "Deny"
    virtual_network_subnet_ids = [var.vm_subnet_id]
  }
}

# Generate a random string for the storage account name
# This ensures globally unique storage account names
resource "random_string" "storage_suffix" {
  length  = 8
  special = false
  upper   = false
}

# Create a blob container for storing data
# Access type is set to private for security
resource "azurerm_storage_container" "blob_container" {
  name                  = "data"
  storage_account_id    = azurerm_storage_account.storage.id
  container_access_type = "private"
}

# Assign Storage Blob Data Contributor role to VMs
# This allows VMs to access the blob storage using their managed identities
resource "azurerm_role_assignment" "vm_contributor" {
  count                = 2
  scope                = azurerm_storage_account.storage.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = var.vm_principal_ids[count.index]
} 