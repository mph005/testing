output "cluster_id" {
  description = "The ID of the AKS cluster"
  value       = module.aks.cluster_id
}

output "kube_config" {
  description = "The kubeconfig to connect to the cluster"
  value       = module.aks.kube_config
  sensitive   = true
}

output "cluster_name" {
  description = "The name of the AKS cluster"
  value       = module.aks.cluster_name
}

output "resource_group_name" {
  description = "The name of the resource group"
  value       = azurerm_resource_group.aks_rg.name
} 