# Azure Kubernetes Service (AKS) Terraform Project

This Terraform project deploys an Azure Kubernetes Service (AKS) cluster using a modular approach and Terraform Cloud for state management.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) (>= 1.0)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) installed and configured
- Azure subscription
- Terraform Cloud account

## Project Structure

```
.
├── main.tf              # Main configuration calling the AKS module
├── provider.tf          # Azure provider and Terraform Cloud configuration
├── variables.tf         # Variable definitions
├── outputs.tf           # Output definitions
├── terraform.tfvars.example  # Example variable values
└── modules/
    └── aks/             # AKS module
        ├── main.tf      # Module resources
        ├── variables.tf # Module variables
        └── outputs.tf   # Module outputs
```

## Terraform Cloud Setup

1. Create a Terraform Cloud account at [app.terraform.io](https://app.terraform.io/) if you don't have one
2. Create a new organization or use an existing one
3. Update the `provider.tf` file with your organization name:
   ```hcl
   cloud {
     organization = "your-organization-name"
     
     workspaces {
       name = "aks-terraform-workspace"
     }
   }
   ```
4. Configure your Azure credentials in Terraform Cloud as environment variables:
   - ARM_CLIENT_ID
   - ARM_CLIENT_SECRET
   - ARM_SUBSCRIPTION_ID
   - ARM_TENANT_ID

## Usage

1. Clone this repository
2. Create a `terraform.tfvars` file using the example as a template:
   ```
   cp terraform.tfvars.example terraform.tfvars
   ```
3. Edit the `terraform.tfvars` file with your specific values
4. Login to Terraform Cloud:
   ```
   terraform login
   ```
5. Initialize Terraform:
   ```
   terraform init
   ```
6. Create a plan:
   ```
   terraform plan
   ```
7. Apply the plan:
   ```
   terraform apply
   ```

The plan and apply operations will be executed on Terraform Cloud. You can monitor the progress in the Terraform Cloud web UI.

## Accessing the Cluster

After deployment, you can configure kubectl to use the new cluster:

```bash
az aks get-credentials --resource-group <resource_group_name> --name <cluster_name>
```

## Cleanup

To destroy all resources created by this project:

```bash
terraform destroy
```