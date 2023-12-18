package main

import (
	"fmt"
	"github.com/username/PluginManager/pkg" // importing the PluginManager package
)

func main() {
	demoData := pkg.NewPluginManagerDemoData()
	data := demoData.GetDemoData()

	for _, v := range data {
		fmt.Printf("Manufacturer: %s, Plugin: %s, Ident: %s\n", v.Manufacturer, v.Plugin, v.Ident)
	}
}
