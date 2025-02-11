variable "storage_account_name" {
  description = "The name of the storage account. It must be between 3 and 24 characters in length and can contain only lowercase letters and numbers."
  type        = string
}

variable "location" {
  description = "The location where the storage account will be created."
  type        = string
}

variable "sku" {
  description = "The SKU of the storage account."
  type        = string
  default     = "Standard_LRS"
}

variable "resource_group_name" {
  description = "The name of the resource group where the storage account will be created."
  type        = string
}

variable "account_tier" {
  description = "The tier of the storage account."
  type        = string
}

variable "account_replication_type" {
  description = "The replication type of the storage account."
  type        = string
}

variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
}

variable "os_disk_name" {
  description = "The name of the OS disk."
  type        = string
}

variable "virtual_machine_name" {
  description = "The name of the virtual machine"
  type        = string
}

variable "vm_size" {
  description = "The size of the virtual machine"
  type        = string
}

variable "os_disk_caching" {
  description = "The caching type for the OS disk"
  type        = string
}

variable "os_disk_create_option" {
  description = "The create option for the OS disk"
  type        = string
}

variable "os_disk_managed_disk_type" {
  description = "The managed disk type for the OS disk"
  type        = string
}

variable "image_publisher" {
  description = "The publisher of the VM image"
  type        = string
}

variable "image_offer" {
  description = "The offer of the VM image"
  type        = string
}

variable "image_sku" {
  description = "The SKU of the VM image"
  type        = string
}

variable "image_version" {
  description = "The version of the VM image"
  type        = string
}

variable "computer_name" {
  description = "The computer name of the VM"
  type        = string
}

variable "admin_username" {
  description = "The admin username for the VM"
  type        = string
}

variable "admin_password" {
  description = "The admin password for the VM"
  type        = string
  sensitive   = true
}