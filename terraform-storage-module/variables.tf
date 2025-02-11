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

variable "storage_account_name" {
  description = "The name of the storage account."
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group."
  type        = string
}

variable "location" {
  description = "The location of the resource group."
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