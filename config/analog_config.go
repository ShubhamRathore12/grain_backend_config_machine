package config

import "strings"

// analogConfigs maps machine name → analog input/output config
var analogConfigs map[string]*AnalogConfig

func init() {
	analogConfigs = make(map[string]*AnalogConfig)

	sharedS7_1200 := &AnalogConfig{
		DisplayName: "S7-1200 Machine",
		Inputs: map[string]string{
			"Suction Pressure":  "LP_value",
			"Discharge Pressure": "HP_value",
			"T2.1 Ambient Temp": "T2_1_ambient_temp",
			"T2.2 Ambient Temp": "T2_2_ambient_temp",
			"T1.1 Cold Temp":    "T1_1_cold_air_temp",
			"T1.2 Cold Temp":    "T1_2_cold_air_temp",
			"T0.1 Air Outlet Temp": "T0_1_air_outlet_temp",
			"T0.2 Air Outlet Temp": "T0_2_air_outlet_temp",
		},
		Outputs: map[string]string{
			"Blower Speed":        "Blower_speed",
			"Condenser fan speed": "Cond_fan_speed",
			"Cond. Fan Speed":     "Condenser_fan_speed",
			"Hot Gas Valve":       "Hot_valve_speed",
			"Afterheat Valve":     "AHT_valve_speed",
		},
	}

	// Assign shared config to all standard S7-1200 machines
	for _, name := range []string{
		"gtpl-122-gt-1000t-s7-1200", "gtpl-121-gt-1000t-s7-1200",
		"gtpl-124-gt-450t-s7-1200",
		"gtpl-133-gt-650t-s7-1200", "gtpl-131-gt-650t-s7-1200",
		"gtpl-081-gt-650t-s7-1200", "gtpl-105-gt-650t-s7-1200",
	} {
		analogConfigs[name] = sharedS7_1200
	}

	// GTPL-132 / AP machines
	gtpl132Config := &AnalogConfig{
		DisplayName: "S7-1200 Machine",
		Inputs: map[string]string{
			"Suction pressure":          "LP_value",
			"Discharge pressure":        "HP_value",
			"T0 probe #1 (Afterheater)": "T0_1_air_outlet_temp",
			"T0 probe #2 (Afterheater)": "T0_2_air_outlet_temp",
			"T1 probe #1 (Cold Air)":    "T1_1_cold_air_temp",
			"T1 probe #2 (Cold Air)":    "T1_2_cold_air_temp",
			"T2 probe #1 (Ambient Air)": "T2_1_ambient_temp",
			"T2 probe #2 (Ambient Air)": "T2_2_ambient_temp",
			"TH probe #1 (Supply Air)":  "TH_1_supply_air_temp",
			"TH probe #2 (Supply Air)":  "TH_2_supply_air_temp",
		},
		Outputs: map[string]string{
			"Blower Speed":        "Blower_speed",
			"Condenser fan speed": "Condenser_fan_speed",
			"Cond. Fan Speed":     "Cond_fan_speed",
			"Hot Gas Valve":       "Hot_valve_speed",
			"Afterheat Valve":     "AHT_valve_speed",
			"Heater":              "Heater_speed",
		},
	}
	analogConfigs["gtpl-132-300-ap-s7-1200"] = gtpl132Config

	// GTPL-134/135 config
	gtpl134Config := &AnalogConfig{
		DisplayName: "S7-1200 Machine",
		Inputs: map[string]string{
			"Suction Pressure":     "LP_value",
			"Discharge Pressure":   "HP_value",
			"T2.1 Ambient Temp":    "T2_temp_mean",
			"T1.1 Cold Temp":       "T1_temp_mean",
			"T0.1 Air Outlet Temp": "T0_temp_mean",
		},
		Outputs: map[string]string{
			"Blower Speed":        "Blower_speed",
			"Condenser fan speed": "Cond_fan_speed",
			"Cond. Fan Speed":     "Condenser_fan_speed",
			"Hot Gas Valve":       "Hot_valve_speed",
			"Afterheat Valve":     "AHT_valve_speed",
		},
	}
	for _, name := range []string{
		"gtpl-134-gt-450t-s7-1200", "gtpl-135-gt-450t-s7-1200",
		"gtpl-145-gt-450t-s7-1200", "gtpl-148-gt-450t-s7-1200",
	} {
		analogConfigs[name] = gtpl134Config
	}

	// GTPL-137/138 config
	gtpl137Config := &AnalogConfig{
		DisplayName: "GTPL-137 Machine",
		Inputs: map[string]string{
			"Suction Pressure":     "LP_value",
			"Discharge Pressure":   "HP_value",
			"T2.1 Ambient Temp":    "T2_1_ambient_temp",
			"T2.2 Ambient Temp":    "T2_2_ambient_temp",
			"T1.1 Cold Temp":       "T1_1_cold_air_temp",
			"T1.2 Cold Temp":       "T1_2_cold_air_temp",
			"T0.1 Air Outlet Temp": "T0_1_air_outlet_temp",
			"T0.2 Air Outlet Temp": "T0_2_air_outlet_temp",
		},
		Outputs: map[string]string{
			"Blower Speed":        "Blower_speed",
			"Condenser fan speed": "Cond_fan_speed",
			"Cond. Fan Speed":     "Condenser_fan_speed",
			"Hot Gas Valve":       "Hot_valve_speed",
			"Afterheat Valve":     "AHT_valve_speed",
		},
	}
	analogConfigs["gtpl-137-gt-450t-s7-1200"] = gtpl137Config
	analogConfigs["gtpl-138-gt-450t-s7-1200"] = gtpl137Config

	// GTPL-061 config
	gtpl061Config := &AnalogConfig{
		DisplayName: "GTPL-061 Machine",
		Inputs: map[string]string{
			"Suction Pressure":     "LP_value",
			"Discharge Pressure":   "HP_value",
			"T2.1 Ambient Temp":    "T2_1_ambient_temp",
			"T2.2 Ambient Temp":    "T2_2_ambient_temp",
			"T1.1 Cold Temp":       "T1_1_cold_air_temp",
			"T1.2 Cold Temp":       "T1_2_cold_air_temp",
			"T0.1 Air Outlet Temp": "T0_1_air_outlet_temp",
			"T0.2 Air Outlet Temp": "T0_2_air_outlet_temp",
		},
		Outputs: map[string]string{
			"Blower Speed":    "Blower_speed",
			"Hot Gas Valve":   "Hot_valve_speed",
			"Afterheat Valve": "AHT_valve_speed",
		},
	}
	analogConfigs["gtpl-061-gt-450t-s7-1200"] = gtpl061Config
}

// GetAnalogConfig returns the analog config for a machine
func GetAnalogConfig(name string) (*AnalogConfig, bool) {
	cfg, ok := analogConfigs[strings.ToLower(name)]
	return cfg, ok
}
