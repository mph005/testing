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

variable "vm_size" {
  description = "Size of the Virtual Machines"
  type        = string
}

variable "admin_username" {
  description = "Username for the VM administrator"
  type        = string
}

variable "admin_password" {
  description = "Password for the VM administrator"
  type        = string
  sensitive   = true
}

variable "subnet_id" {
  description = "ID of the subnet where VMs will be created"
  type        = string
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
} 