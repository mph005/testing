# Terraform Storage Module

This Terraform module is designed to create and manage Azure Storage Accounts. It provides a simple way to define and deploy storage accounts with customizable settings.

## Overview

This module allows users to create Azure Storage Accounts with specified configurations such as name, location, and SKU. It abstracts the complexity of the underlying Terraform configurations, making it easier to deploy storage resources.

## Usage

To use this module, include it in your Terraform configuration as follows:

```hcl
module "storage" {
  source              = "./terraform-storage-module"
  storage_account_name = "your_storage_account_name"
  location           = "East US"
  sku                = "Standard_LRS"
}
```

## Inputs

| Name                    | Description                                  | Type   | Default | Required |
|-------------------------|----------------------------------------------|--------|---------|:--------:|
| storage_account_name    | The name of the storage account.             | string | n/a     | yes      |
| location                | The location where the storage account will be created. | string | n/a     | yes      |
| sku                     | The SKU of the storage account.              | string | "Standard_LRS" | no       |

## Outputs

| Name                    | Description                                  |
|-------------------------|----------------------------------------------|
| storage_account_id      | The ID of the created storage account.       |
| primary_endpoint        | The primary endpoint of the storage account. |

## Example

Here is an example of how to use this module in a Terraform configuration:

```hcl
provider "azurerm" {
  features {}
}

module "storage" {
  source              = "./terraform-storage-module"
  storage_account_name = "examplestoracc"
  location           = "East US"
  sku                = "Standard_LRS"
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.