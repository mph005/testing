variable "resource_group_name" {
  description = "Name of the resource group to create"
  type        = string
}

variable "resource_group_location" {
  description = "Location for the resource group"
  type        = string
  default     = "East US"
}

variable "cluster_name" {
  description = "Name of the AKS cluster"
  type        = string
}

variable "kubernetes_version" {
  description = "Kubernetes version to use"
  type        = string
  default     = "1.27"
}

variable "node_count" {
  description = "Number of nodes in the default node pool"
  type        = number
  default     = 2
}

variable "vm_size" {
  description = "Size of the VM for the nodes"
  type        = string
  default     = "Standard_D2s_v3"
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default     = {}
} 