package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/Stelminator/terraform-provider-launchdarkly/launchdarkly"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: launchdarkly.Provider})
}
