resource "azurerm_availability_set" "vm_avset" {
  name                         = "avset-${var.environment}"
  location                     = var.location
  resource_group_name          = var.resource_group_name
  platform_fault_domain_count  = 2
  platform_update_domain_count = 5
  tags                         = var.tags
}

resource "azurerm_network_interface" "vm_nic" {
  count               = 2
  name                = "nic-vm${count.index + 1}-${var.environment}"
  location            = var.location
  resource_group_name = var.resource_group_name
  tags                = var.tags

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_linux_virtual_machine" "vm" {
  count                           = 2
  name                            = "vm${count.index + 1}-${var.environment}"
  resource_group_name             = var.resource_group_name
  location                        = var.location
  size                           = var.vm_size
  admin_username                  = var.admin_username
  admin_password                  = var.admin_password
  disable_password_authentication = false
  availability_set_id            = azurerm_availability_set.vm_avset.id
  network_interface_ids          = [azurerm_network_interface.vm_nic[count.index].id]
  tags                           = var.tags

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_virtual_machine_extension" "blob_extension" {
  count                = 2
  name                 = "blobfuse-extension"
  virtual_machine_id   = azurerm_linux_virtual_machine.vm[count.index].id
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.1"

  settings = jsonencode({
    "commandToExecute" = "apt-get update && apt-get install -y blobfuse2"
  })

  tags = var.tags
} 