# FILE: /terraform-storage-project/terraform-storage-project/README.md

# Terraform Azure Storage Project

This project utilizes Terraform to create an Azure Storage Account using a dedicated module. The configuration is designed to be reusable and modular, allowing for easy adjustments and scalability.

## Project Structure

- `main.tf`: Contains the main Terraform configuration that calls the `terraform-storage-module` to create a new Azure storage account.
- `variables.tf`: Defines the input variables for the Terraform configuration, including storage account details and settings.
- `outputs.tf`: Specifies the output values that will be displayed after the deployment, such as the storage account ID and primary endpoint.
- `terraform-storage-module/`: A subdirectory containing the module for creating Azure resources.
  - `main.tf`: Resource definitions for the storage account and related resources.
  - `variables.tf`: Input variables specific to the module.
  - `outputs.tf`: Output values for the module.
  - `providers.tf`: Required providers for the module, including Azure provider configuration.
  - `README.md`: Documentation for the module, including usage instructions and examples.

## Getting Started

### Prerequisites

- Terraform installed on your machine.
- An Azure account with appropriate permissions to create resources.

### Initialization

To initialize the Terraform configuration, navigate to the project directory and run:

```
terraform init
```

### Planning

To see what changes will be made by Terraform, run:

```
terraform plan
```

### Applying

To apply the configuration and create the resources, run:

```
terraform apply
```

### Outputs

After the deployment is complete, Terraform will display the output values defined in `outputs.tf`, including the storage account ID and primary endpoint.

## Usage

You can customize the variables defined in `variables.tf` to suit your needs. Make sure to provide valid values for the required variables before applying the configuration.

## License

This project is licensed under the MIT License.