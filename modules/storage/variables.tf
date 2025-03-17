variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
}

variable "location" {
  description = "Azure region for resources"
  type        = string
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "storage_account_tier" {
  description = "Storage account tier"
  type        = string
}

variable "storage_account_replication_type" {
  description = "Storage account replication type"
  type        = string
}

variable "vm_subnet_id" {
  description = "ID of the VM subnet"
  type        = string
}

variable "vm_principal_ids" {
  description = "List of VM principal IDs for role assignment"
  type        = list(string)
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
} 