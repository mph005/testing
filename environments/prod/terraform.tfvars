resource_group_name             = "rg-vm-storage-prod"
location                      = "centralus"
environment                   = "prod"
vm_size                      = "Standard_D2s_v3"
vm_admin_username            = "azureuser"
storage_account_tier         = "Standard"
storage_account_replication_type = "LRS"

tags = {
  Environment = "Production"
  ManagedBy   = "Terraform"
  Project     = "VM Storage Demo"
} 