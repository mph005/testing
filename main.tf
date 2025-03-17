resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
  tags     = var.tags
}

module "networking" {
  source              = "./modules/networking"
  resource_group_name = azurerm_resource_group.rg.name
  location           = var.location
  environment        = var.environment
  tags               = var.tags
}

module "storage" {
  source                        = "./modules/storage"
  resource_group_name           = azurerm_resource_group.rg.name
  location                      = var.location
  environment                   = var.environment
  storage_account_tier          = var.storage_account_tier
  storage_account_replication_type = var.storage_account_replication_type
  vm_subnet_id                  = module.networking.vm_subnet_id
  vm_principal_ids             = module.compute.vm_principal_ids
  tags                         = var.tags
}

module "compute" {
  source              = "./modules/compute"
  resource_group_name = azurerm_resource_group.rg.name
  location           = var.location
  environment        = var.environment
  vm_size           = var.vm_size
  admin_username     = var.vm_admin_username
  admin_password     = var.vm_admin_password
  subnet_id         = module.networking.vm_subnet_id
  tags              = var.tags
} 