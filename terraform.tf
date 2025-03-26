# Terraform configuration file
# This file defines the required providers, versions, and backend configuration

# Specify the required Terraform version
# Using version 1.0.0 or higher for stability and features
terraform {
  required_version = ">= 1.0.0"

  # Configure required providers with version constraints
  # Using Azure RM provider version 3.0 or higher
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }

  # Configure Terraform Cloud as the backend
  # This provides state management, version control, and collaboration features
  cloud {
    organization = "your-org-name"
    workspaces {
      name = "azure-vm-storage"
    }
  }
}

# Configure the Azure provider with custom features
# Enable resource group protection to prevent accidental deletion
provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = true
    }
  }
} 