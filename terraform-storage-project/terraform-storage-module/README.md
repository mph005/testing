# FILE: /terraform-storage-project/terraform-storage-project/terraform-storage-module/README.md

# Terraform Storage Module

This module creates an Azure Storage Account along with a Resource Group. It allows you to specify various parameters for the storage account and resource group.

## Usage

To use this module, include it in your Terraform configuration as follows:

```hcl
module "storage_account" {
  source                  = "./terraform-storage-module"
  storage_account_name    = var.storage_account_name
  resource_group_name     = var.resource_group_name
  location                = var.location
  account_tier            = var.account_tier
  account_replication_type = var.account_replication_type
  tags                    = var.tags
}
```

## Inputs

| Name                       | Description                                         | Type   | Default | Required |
|----------------------------|-----------------------------------------------------|--------|---------|:--------:|
| storage_account_name       | The name of the storage account.                    | string | n/a     |   yes    |
| resource_group_name        | The name of the resource group.                     | string | n/a     |   yes    |
| location                   | The Azure region where the resources will be created. | string | n/a     |   yes    |
| account_tier               | The tier of the storage account (e.g., Standard, Premium). | string | Standard | no     |
| account_replication_type   | The replication type for the storage account (e.g., LRS, GRS). | string | LRS     | no     |
| tags                       | A map of tags to assign to the resources.          | map    | {}      | no      |

## Outputs

| Name                  | Description                                   |
|-----------------------|-----------------------------------------------|
| storage_account_id    | The ID of the created storage account.        |
| primary_blob_endpoint | The primary blob service endpoint.            |

## Example

```hcl
module "storage_account" {
  source                  = "./terraform-storage-module"
  storage_account_name    = "examplestoracc"
  resource_group_name     = "example-rg"
  location                = "East US"
  account_tier            = "Standard"
  account_replication_type = "LRS"
  tags                    = {
    Environment = "Dev"
    Project     = "Terraform"
  }
}
```

## Requirements

- Terraform 0.12 or later
- Azure Provider

## Author

Your Name

## License

This module is licensed under the MIT License.