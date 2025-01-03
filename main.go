package main

import (
	"fmt"
	"os"

	sbcBakeryVersion "github.com/bryborge/sbc-bakery/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	// Register Builders, Provisioners, Post-Processors, and Datasources here ...
	pps.SetVersion(sbcBakeryVersion.PluginVersion)

	err := pps.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
