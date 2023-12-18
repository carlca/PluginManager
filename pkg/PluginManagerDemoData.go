package PluginManager

import (
	"strings"
)

type PluginManagerDemoData struct {
	DemoData []PluginTriplet
}

func (p *PluginManagerDemoData) BuildTriplet(csvLine string) []PluginTriplet {
	parts := strings.Split(csvLine, ",")
	if len(parts) == 3 {
		return []PluginTriplet{{Manufacturer: parts[0], Plugin: parts[1], Ident: parts[2]}}
	}
	return []PluginTriplet{{Manufacturer: "invalid", Plugin: "invalid", Ident: "invalid"}}
}

func (p *PluginManagerDemoData) GetDemoData() []PluginTriplet {
	lines := []string{
		"audio,Vital.clap,audio.vital.synth",
		"audiodamage,Dubstation2.clap,com.audiodamage.dubstation2",
		"audiomodern,Chordjam.clap,com.audiomodern.chordjam",
		// Add more CSV lines here...
	}

	for _, line := range lines {
		p.DemoData = append(p.DemoData, p.BuildTriplet(line)...)
	}

	return p.DemoData
}

func Run() {
	demoData := PluginManagerDemoData{}
	data := demoData.GetDemoData()

	// print data for testing
	for _, v := range data {
		println(v.Manufacturer, v.Plugin, v.Ident)
	}
}
