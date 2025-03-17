variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
  default     = "rg-vm-storage-prod"
}

variable "location" {
  description = "Azure region for resources"
  type        = string
  default     = "centralus"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "prod"
}

variable "vm_size" {
  description = "Size of the Virtual Machines"
  type        = string
  default     = "Standard_D2s_v3"
}

variable "vm_admin_username" {
  description = "Username for the VM administrator"
  type        = string
  default     = "azureuser"
}

variable "vm_admin_password" {
  description = "Password for the VM administrator"
  type        = string
  sensitive   = true
}

variable "storage_account_tier" {
  description = "Storage account tier"
  type        = string
  default     = "Standard"
}

variable "storage_account_replication_type" {
  description = "Storage account replication type"
  type        = string
  default     = "LRS"
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default = {
    Environment = "Production"
    ManagedBy   = "Terraform"
    Project     = "VM Storage Demo"
  }
} 