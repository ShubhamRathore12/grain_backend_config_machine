package config

import "strings"

// autoConfigs maps machine name → auto mode config
var autoConfigs = map[string]*AutoConfig{
	"gtpl-122-gt-1000t-s7-1200": {
		SerialNumber: "GTPL_122_S7_1200",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-121-gt-1000t-s7-1200": {
		SerialNumber: "GTPL_121",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-081-gt-650t-s7-1200": {
		SerialNumber: "GTPL_081",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-105-gt-650t-s7-1200": {
		SerialNumber: "GTPL_105",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-133-gt-650t-s7-1200": {
		SerialNumber: "GTPL_132",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-131-gt-650t-s7-1200": {
		SerialNumber: "GTPL_131",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-124-gt-450t-s7-1200": {
		SerialNumber: "GTPL_124",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-137-gt-450t-s7-1200": {
		SerialNumber: "GTPL_137",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
			"COND":   {Key: "Cond_fan_speed", Label: "Condenser Fan"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-138-gt-450t-s7-1200": {
		SerialNumber: "GTPL_138",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
			"COND":   {Key: "Cond_fan_speed", Label: "Condenser Fan"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-134-gt-450t-s7-1200": {
		SerialNumber: "GTPL_134",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-135-gt-450t-s7-1200": {
		SerialNumber: "GTPL_135",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-145-gt-450t-s7-1200": {
		SerialNumber: "GTPL_145",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-148-gt-450t-s7-1200": {
		SerialNumber: "GTPL_148",
		TemperatureSensors: map[string]SensorConfig{
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":    {Key: "AHT_valve_speed", Label: "After Heat(AHT)"},
			"HGS":    {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER": {Key: "Blower_speed", Label: "Blower"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-30-gt-180e-s7-1200": {
		SerialNumber: "GTPL_114",
		TemperatureSensors: map[string]SensorConfig{
			"TH": {Key: "TH_temp_mean", Label: "Supply Air(TH)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":             {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":             {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER":          {Key: "Blower_speed", Label: "Blower"},
			"CONDENSORFANSPEED": {Key: "Cond_fan_speed", Label: "Cond. Fan Speed"},
			"HTR":             {Key: "Heater_speed", Label: "Heater"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
	"gtpl-115-gt-180e-s7-1200": {
		SerialNumber: "GTPL_115",
		TemperatureSensors: map[string]SensorConfig{
			"TH": {Key: "TH_temp_mean", Label: "Supply Air(TH)"},
			"T1": {Key: "T1_temp_mean", Label: "Cold Air(T1)"},
			"T2": {Key: "T2_temp_mean", Label: "Ambient(T2)"},
			"T0": {Key: "T0_temp_mean", Label: "Air Outlet(T0)"},
		},
		Controls: map[string]ControlConfig{
			"AHT":             {Key: "AHT_vale_speed", Label: "After Heat(AHT)"},
			"HGS":             {Key: "Hot_valve_speed", Label: "Hot Gas(HGS)"},
			"BLOWER":          {Key: "Blower_speed", Label: "Blower"},
			"CONDENSORFANSPEED": {Key: "Cond_fan_speed", Label: "Cond. Fan Speed"},
			"HTR":             {Key: "Heater_speed", Label: "Heater"},
		},
		Compressor: CompressorConfig{Time: "Compressor_timer", HP: "HP_value", LP: "LP_value"},
	},
}

// GetAutoConfig returns the auto mode config for a machine
func GetAutoConfig(name string) (*AutoConfig, bool) {
	cfg, ok := autoConfigs[strings.ToLower(name)]
	return cfg, ok
}

// GetGrainConfig returns the grain mode config (same structure, different serial)
func GetGrainConfig(name string) (*AutoConfig, bool) {
	key := strings.ToLower(name)
	base, ok := autoConfigs[key]
	if !ok {
		return nil, false
	}
	// Grain config uses same sensors/controls but different serial number
	grain := *base
	grain.SerialNumber = base.SerialNumber + "_GRAIN"
	return &grain, true
}

// GetPaddyConfig returns the paddy mode config
func GetPaddyConfig(name string) (*AutoConfig, bool) {
	key := strings.ToLower(name)
	base, ok := autoConfigs[key]
	if !ok {
		return nil, false
	}
	paddy := *base
	paddy.SerialNumber = base.SerialNumber + "_PADDY"
	return &paddy, true
}
