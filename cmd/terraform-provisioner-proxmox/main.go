package main

import (
	"github.com/Telmate/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: func() terraform.ResourceProvisioner {
			return proxmox.Provisioner()
		},
	})
}
