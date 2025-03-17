output "cluster_id" {
  description = "The ID of the AKS cluster"
  value       = azurerm_kubernetes_cluster.aks.id
}

output "cluster_name" {
  description = "The name of the AKS cluster"
  value       = azurerm_kubernetes_cluster.aks.name
}

output "kube_config" {
  description = "The kubeconfig to connect to the cluster"
  value       = azurerm_kubernetes_cluster.aks.kube_config_raw
  sensitive   = true
}

output "kube_config_host" {
  description = "The Kubernetes cluster server host"
  value       = azurerm_kubernetes_cluster.aks.kube_config.0.host
}

output "client_certificate" {
  description = "The client certificate for authentication"
  value       = azurerm_kubernetes_cluster.aks.kube_config.0.client_certificate
  sensitive   = true
} 