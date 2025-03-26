# Networking module for Azure VM Storage Project
# This module creates the network infrastructure including VNet, subnets, NSGs, and Bastion host

# Create the Virtual Network that will host all resources
# Using a /16 address space to allow for future expansion
resource "azurerm_virtual_network" "vnet" {
  name                = "vnet-${var.environment}"
  address_space       = ["10.0.0.0/16"]
  location            = var.location
  resource_group_name = var.resource_group_name
  tags                = var.tags
}

# Create a subnet for the VMs
# Using a /24 address space which allows for 254 usable IP addresses
resource "azurerm_subnet" "vm_subnet" {
  name                 = "snet-vm"
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}

# Create the Azure Bastion subnet
# This is a dedicated subnet required for Azure Bastion host
resource "azurerm_subnet" "bastion_subnet" {
  name                 = "AzureBastionSubnet"
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}

# Create a public IP for the Bastion host
# This is required for secure VM access through Azure Bastion
resource "azurerm_public_ip" "bastion_ip" {
  name                = "pip-bastion"
  location            = var.location
  resource_group_name = var.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"
  tags                = var.tags
}

# Create the Azure Bastion host
# This provides secure access to the VMs without exposing them to the internet
resource "azurerm_bastion_host" "bastion" {
  name                = "bastion-host"
  location            = var.location
  resource_group_name = var.resource_group_name
  tags                = var.tags

  ip_configuration {
    name                 = "configuration"
    subnet_id            = azurerm_subnet.bastion_subnet.id
    public_ip_address_id = azurerm_public_ip.bastion_ip.id
  }
}

# Create Network Security Group for the VMs
# This defines the security rules for VM access
resource "azurerm_network_security_group" "vm_nsg" {
  name                = "nsg-vm"
  location            = var.location
  resource_group_name = var.resource_group_name
  tags                = var.tags

  # Allow inbound SSH (port 22) and RDP (port 3389) only from the Virtual Network
  # This ensures VMs are only accessible through Azure Bastion
  security_rule {
    name                       = "AllowBastionInbound"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_ranges    = ["22", "3389"]
    source_address_prefix      = "VirtualNetwork"
    destination_address_prefix = "*"
  }
}

# Associate the NSG with the VM subnet
# This applies the security rules to all VMs in the subnet
resource "azurerm_subnet_network_security_group_association" "vm_nsg_association" {
  subnet_id                 = azurerm_subnet.vm_subnet.id
  network_security_group_id = azurerm_network_security_group.vm_nsg.id
} 