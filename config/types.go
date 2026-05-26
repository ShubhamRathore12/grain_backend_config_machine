package config

// Machine represents the full configuration for a single machine
type Machine struct {
	Name           string          `json:"name"`
	Location       string          `json:"location"`
	Image          string          `json:"image"`
	PLC            string          `json:"plc"`
	ChillerModel   string          `json:"chillerModel"`
	Table          string          `json:"table"`
	StatusKey      string          `json:"statusKey"`
	Tags           []string        `json:"tags"`
	MenuType       string          `json:"menuType"`       // "standard", "grainPaddy"
	AerationType   string          `json:"aerationType"`   // "withHeating", "withoutHeating"
	PressureUnit   string          `json:"pressureUnit"`   // "psi", "bar"
	CondenserFans  int             `json:"condenserFans"`  // 1, 2, 4, 6
	HasHTR         bool            `json:"hasHTR"`         // has heater in diagram
	HasAnalogMenu  bool            `json:"hasAnalogMenu"`
	IsS7200        bool            `json:"isS7200"`
	AutoConfig     *AutoConfig     `json:"autoConfig,omitempty"`
	GrainConfig    *AutoConfig     `json:"grainConfig,omitempty"`
	PaddyConfig    *AutoConfig     `json:"paddyConfig,omitempty"`
	OutputsConfig  []OutputItem    `json:"outputsConfig,omitempty"`
	AnalogConfig   *AnalogConfig   `json:"analogConfig,omitempty"`
}

// AutoConfig holds temperature sensors, controls, and compressor config
type AutoConfig struct {
	SerialNumber       string                    `json:"serialNumber"`
	TemperatureSensors map[string]SensorConfig   `json:"temperatureSensors"`
	Controls           map[string]ControlConfig  `json:"controls"`
	Compressor         CompressorConfig          `json:"compressor"`
}

// SensorConfig maps a sensor key to its data field and label
type SensorConfig struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// ControlConfig maps a control to its data field and label
type ControlConfig struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// CompressorConfig holds compressor-related field mappings
type CompressorConfig struct {
	Time string `json:"time"`
	HP   string `json:"hp"`
	LP   string `json:"lp"`
}

// OutputItem represents a single output pin configuration
type OutputItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DataKey     string `json:"dataKey"`
}

// AnalogConfig holds analog input/output field mappings
type AnalogConfig struct {
	DisplayName string            `json:"displayName"`
	Inputs      map[string]string `json:"inputs"`
	Outputs     map[string]string `json:"outputs"`
}

// MenuConfig describes what menu items a machine should show
type MenuConfig struct {
	MenuType      string     `json:"menuType"`
	Items         []MenuItem `json:"items"`
	ShowAnalog    bool       `json:"showAnalog"`
}

// MenuItem represents a single menu entry
type MenuItem struct {
	Title    string `json:"title"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	Gradient string `json:"gradient"`
}
