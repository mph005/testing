# Azure VM Storage Project

This Terraform project deploys two Ubuntu VMs with shared blob storage in Azure. The infrastructure includes:

- Two Ubuntu 22.04 LTS VMs in an availability set
- Azure Blob Storage with private endpoint
- Azure Bastion for secure VM access
- Network security groups and VNet configuration

## Prerequisites

- Terraform 1.0.0 or newer
- Azure subscription
- Terraform Cloud account
- Azure CLI installed locally

## Project Structure

```
.
├── README.md
├── main.tf
├── variables.tf
├── outputs.tf
├── terraform.tf
├── modules/
│   ├── networking/
│   ├── compute/
│   └── storage/
└── environments/
    └── prod/
        └── terraform.tfvars
```

## Configuration

1. Update the Terraform Cloud organization name in `terraform.tf`
2. Create a workspace in Terraform Cloud
3. Configure the following environment variables in Terraform Cloud:
   - `ARM_SUBSCRIPTION_ID`
   - `ARM_TENANT_ID`
   - `ARM_CLIENT_ID`
   - `ARM_CLIENT_SECRET`
4. Update the `vm_admin_password` variable (do not store in version control)

## Usage

1. Initialize Terraform:
   ```bash
   terraform init
   ```

2. Review the plan:
   ```bash
   terraform plan
   ```

3. Apply the configuration:
   ```bash
   terraform apply
   ```

## Accessing the VMs

The VMs are accessible only through Azure Bastion. Use the following steps:

1. Connect to the Azure Portal
2. Navigate to the VM resource
3. Click "Connect" and select "Bastion"
4. Enter the username and password configured in the variables

## Mounting the Blob Storage

The VMs are pre-configured with blobfuse2. To mount the blob storage:

1. Create a mount point:
   ```bash
   sudo mkdir /mnt/blob
   ```

2. Configure blobfuse2 using the storage account credentials (available in Azure Portal)

## Security Features

- No public IP addresses on VMs
- Azure Bastion for secure access
- Network security groups with minimal required access
- Private endpoint for blob storage
- System-assigned managed identities for VMs

## Tags

Resources are tagged with:
- Environment
- ManagedBy
- Project

## Contributing

1. Fork the repository
2. Create a feature branch
3. Submit a pull request 