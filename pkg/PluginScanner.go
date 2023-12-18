package PluginManager

type PluginScanner struct {
	PlugPlugins []PluginTriplet
	// other fields...
}

func NewPluginScanner() *PluginScanner {
	return &PluginScanner{
		PlugPlugins: []PluginTriplet{},
		// initialize other fields...
	}
}
