package eiscp

// These structs are used for unmarshaling the XML retrieved from the receiver

// Onkyo XML info
type onkyoXML struct {
	Device onkyoXMLDeviceInfo `xml:"device"`
}

// Device Info (DEVICE)
type onkyoXMLDeviceInfo struct {
	Brand        string         `xml:"brand"`        // Brand (Onkyo, Integra, Pioneer?)
	ModelName    string         `xml:"model"`        // Model number
	FriendlyName string         `xml:"friendlyname"` // Friendly name (user-definable at web interface)
	PresetList   onkyoXMLPreset `xml:"presetlist"`   // Tuner presets
	InputList    onkyoXMLInput  `xml:"selectorlist"` // Enabled sources
}

// Preset list (PRESETLIST items)
type onkyoXMLPreset struct {
	Preset []onkyoXMLPresetItem `xml:"preset"`
}

// Input list (SELECTORLIST items)
type onkyoXMLInput struct {
	Input []onkyoXMLInputItem `xml:"selector"`
}

// Tuner preset
type onkyoXMLPresetItem struct {
	Id        string `xml:"id,attr"`   // Hexadecimal id
	Frequency string `xml:"freq,attr"` // Frequency
	Band      int    `xml:"band,attr"` // Band (0: Not Set, 1: FM, 2: AM)
}

// Input declaration
type onkyoXMLInputItem struct {
	Id      string `xml:"id,attr"`    // hexadecimal id of the input
	Enabled bool   `xml:"value,attr"` // Input enabled
	Name    string `xml:"name,attr"`  // input name
}
