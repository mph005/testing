resource "azurerm_resource_group" "aks_rg" {
  name     = var.resource_group_name
  location = var.resource_group_location
  tags     = var.tags
}

module "aks" {
  source = "./modules/aks"

  resource_group_name     = azurerm_resource_group.aks_rg.name
  resource_group_location = azurerm_resource_group.aks_rg.location
  cluster_name            = var.cluster_name
  kubernetes_version      = var.kubernetes_version
  node_count              = var.node_count
  vm_size                 = var.vm_size
  tags                    = var.tags
} 