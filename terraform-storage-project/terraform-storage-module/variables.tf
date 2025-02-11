variable "storage_account_name" {
  description = "The name of the storage account."
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group."
  type        = string
}

variable "location" {
  description = "The Azure location where the resources will be created."
  type        = string
}

variable "account_tier" {
  description = "The tier of the storage account (e.g., Standard, Premium)."
  type        = string
}

variable "account_replication_type" {
  description = "The replication type of the storage account (e.g., LRS, GRS)."
  type        = string
}

variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
}