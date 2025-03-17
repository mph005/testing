terraform {
  cloud {
    organization = "your-organization-name"
    
    workspaces {
      name = "aks-terraform-workspace"
    }
  }

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
  required_version = ">= 1.0"
}

provider "azurerm" {
  features {}
} 