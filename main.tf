# Main Terraform configuration file for Azure VM Storage Project
# This file orchestrates the creation of all infrastructure components by calling the respective modules

# Create the main resource group that will contain all other resources
resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
  tags     = var.tags
}

# Deploy the networking infrastructure including VNet, subnets, NSGs, and Bastion host
# This module sets up the network foundation for the VMs and storage
module "networking" {
  source              = "./modules/networking"
  resource_group_name = azurerm_resource_group.rg.name
  location            = var.location
  environment         = var.environment
  tags                = var.tags
}

# Deploy the storage infrastructure including storage account and blob container
# This module creates the shared storage that will be accessible by the VMs
module "storage" {
  source                           = "./modules/storage"
  resource_group_name              = azurerm_resource_group.rg.name
  location                         = var.location
  environment                      = var.environment
  storage_account_tier             = var.storage_account_tier
  storage_account_replication_type = var.storage_account_replication_type
  vm_subnet_id                     = module.networking.vm_subnet_id
  vm_principal_ids                 = module.compute.vm_principal_ids
  tags                             = var.tags
}

# Deploy the compute infrastructure including VMs and availability set
# This module creates the VMs that will access the shared storage
module "compute" {
  source              = "./modules/compute"
  resource_group_name = azurerm_resource_group.rg.name
  location            = var.location
  environment         = var.environment
  vm_size             = var.vm_size
  admin_username      = var.vm_admin_username
  admin_password      = var.vm_admin_password
  subnet_id           = module.networking.vm_subnet_id
  tags                = var.tags
} 