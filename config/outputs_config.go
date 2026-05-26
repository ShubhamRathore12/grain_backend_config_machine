package config

import "strings"

// outputsConfigs maps machine name → output pin configurations
var outputsConfigs map[string][]OutputItem

func init() {
	outputsConfigs = make(map[string][]OutputItem)

	// 6-fan machines (GTPL-122, 121, 133, 131, 081, 105)
	sixFanOutputs := []OutputItem{
		{ID: "Q0.0", Description: "Compressor Start", DataKey: "Compressor_start_Q0_0"},
		{ID: "Q0.1", Description: "Compressor Module Reset", DataKey: "Compressor_module_reset_Q0_1"},
		{ID: "Q0.2", Description: "CR Valve 25% ON", DataKey: "CR_valve_25_percent_ON_Q0_2"},
		{ID: "Q0.3", Description: "CR Valve 50% ON", DataKey: "CR_valve_50_percent_ON_Q0_3"},
		{ID: "Q0.4", Description: "Solenoid Valve ON", DataKey: "Solenoid_valve_ON_Q0_4"},
		{ID: "Q0.5", Description: "Hot Gas Valve ON", DataKey: "Hot_gas_valve_ON_Q0_5"},
		{ID: "Q0.6", Description: "AHT Valve ON", DataKey: "AHT_valve_ON_Q0_6"},
		{ID: "Q0.7", Description: "Blower Drive Start", DataKey: "Blower_drive_start_Q0_7"},
		{ID: "Q1.0", Description: "System Warning", DataKey: "System_warning_Q1_0"},
		{ID: "Q1.1", Description: "Chiller Healthy", DataKey: "Chiller_healthy_Q1_1"},
		{ID: "Q2.1", Description: "Cond Fan 1 ON", DataKey: "Cond_fan_1_ON_Q2_1"},
		{ID: "Q2.2", Description: "CR Valve 75% ON", DataKey: "CR_valve_75_percent_ON_Q2_2"},
		{ID: "Q2.3", Description: "Chiller Fault", DataKey: "Chiller_fault_Q2_3"},
		{ID: "Q2.4", Description: "Cond Fan 2 ON", DataKey: "Cond_fan_2_ON_Q2_4"},
		{ID: "Q2.5", Description: "Cond Fan 3 ON", DataKey: "Cond_fan_3_ON_Q2_5"},
		{ID: "Q2.6", Description: "Cond Fan 4 ON", DataKey: "Cond_fan_4_ON_Q2_6"},
		{ID: "Q2.7", Description: "CR Valve 100% ON", DataKey: "CR_valve_100_percent_ON_Q2_7"},
		{ID: "Q3.0", Description: "Cond Fan 5 ON", DataKey: "Cond_fan_5_ON_Q3_0"},
		{ID: "Q3.1", Description: "Cond Fan 6 ON", DataKey: "Cond_fan_6_ON_Q3_1"},
	}
	for _, name := range []string{
		"gtpl-122-gt-1000t-s7-1200", "gtpl-121-gt-1000t-s7-1200",
		"gtpl-133-gt-650t-s7-1200", "gtpl-131-gt-650t-s7-1200",
		"gtpl-081-gt-650t-s7-1200", "gtpl-105-gt-650t-s7-1200",
	} {
		outputsConfigs[name] = sixFanOutputs
	}

	// GTPL-118 (S7-200 small)
	outputsConfigs["gtpl-118-gt-60t-s7-200"] = []OutputItem{
		{ID: "Q0.0", Description: "Compressor on", DataKey: "Compressor_on_Q0_0"},
		{ID: "Q0.1", Description: "Compressor motor reset", DataKey: "Compressor_motor_reset_Q0_1"},
		{ID: "Q0.2", Description: "Solenoid valve on", DataKey: "Solenoid_valve_on_Q0_2"},
		{ID: "Q0.3", Description: "Hot gas valve on", DataKey: "Hot_gas_valve_on_Q0_3"},
		{ID: "Q0.4", Description: "After heat valve on", DataKey: "After_heat_valve_on_Q0_4"},
		{ID: "Q0.5", Description: "Blower drive on", DataKey: "Blower_drive_on_Q0_5"},
		{ID: "Q0.6", Description: "Collective trouble signal", DataKey: "Collective_Trouble_Signal_Q0_6"},
		{ID: "Q0.7", Description: "Chiller healthy", DataKey: "Chiller_healthy_on_Q0_7"},
		{ID: "Q1.0", Description: "Condenser fan on", DataKey: "Condenser_fan_on_Q1_0"},
		{ID: "Q1.1", Description: "Chiller fault", DataKey: "Chiller_Fault_Q1_1"},
	}

	// GTPL-134/135/145/148 (4-fan with CR valves)
	fourFanCROutputs := []OutputItem{
		{ID: "Q0.0", Description: "Compressor on", DataKey: "Compressor_on_Q0_0"},
		{ID: "Q0.1", Description: "Compressor motor reset", DataKey: "Compressor_motor_reset_Q0_1"},
		{ID: "Q0.2", Description: "CR 25% ON", DataKey: "CR_25_percent_ON_Q0_2"},
		{ID: "Q0.3", Description: "CR 50% ON", DataKey: "CR_50_percent_ON_Q0_3"},
		{ID: "Q0.4", Description: "Solenoid valve on", DataKey: "Solenoid_valve_on_Q0_4"},
		{ID: "Q0.5", Description: "Hot gas valve on", DataKey: "Hot_gas_valve_on_Q0_5"},
		{ID: "Q0.6", Description: "After heat valve on", DataKey: "After_heat_valve_on_Q0_6"},
		{ID: "Q0.7", Description: "Blower drive on", DataKey: "Blower_drive_on_Q0_7"},
		{ID: "Q1.0", Description: "Collective trouble signal", DataKey: "Collective_trouble_signal_Q1_0"},
		{ID: "Q1.1", Description: "Chiller healthy on", DataKey: "Chiller_healthy_on_Q1_1"},
		{ID: "Q2.0", Description: "Spare", DataKey: "Spare_Q2_0"},
		{ID: "Q2.1", Description: "Condenser fan1 on", DataKey: "Condenser_fan1_on_Q2_1"},
		{ID: "Q2.2", Description: "CR valve 75% on", DataKey: "CR_valve_75_percent_on_Q2_2"},
		{ID: "Q2.3", Description: "Chiller fault", DataKey: "Chiller_fault_Q2_3"},
		{ID: "Q2.4", Description: "Condenser fan2 on", DataKey: "Condenser_fan2_on_Q2_4"},
		{ID: "Q2.5", Description: "Condenser fan3 on", DataKey: "Condenser_fan3_on_Q2_5"},
		{ID: "Q2.6", Description: "Condenser fan4 on", DataKey: "Condenser_fan4_on_Q2_6"},
		{ID: "Q2.7", Description: "CR 100% ON", DataKey: "CR_100_percent_ON_Q2_7"},
	}
	for _, name := range []string{
		"gtpl-134-gt-450t-s7-1200", "gtpl-135-gt-450t-s7-1200",
		"gtpl-145-gt-450t-s7-1200", "gtpl-148-gt-450t-s7-1200",
	} {
		outputsConfigs[name] = fourFanCROutputs
	}

	// GTPL-137/138 (Thailand machines)
	thailandOutputs := []OutputItem{
		{ID: "Q0.0", Description: "Compressor", DataKey: "Compressor_on_Q0_0"},
		{ID: "Q0.1", Description: "Compressor motor reset", DataKey: "Compressor_motor_reset_Q0_1"},
		{ID: "Q0.2", Description: "Spare", DataKey: "Spare_Q0_2"},
		{ID: "Q0.3", Description: "Spare", DataKey: "Spare_Q0_3"},
		{ID: "Q0.4", Description: "Solenoid valve on", DataKey: "Solenoid_valve_on_Q0_4"},
		{ID: "Q0.5", Description: "Hot gas valve on", DataKey: "Hot_gas_valve_on_Q0_5"},
		{ID: "Q0.6", Description: "After heat valve on", DataKey: "After_heat_valve_on_Q0_6"},
		{ID: "Q0.7", Description: "Blower drive on", DataKey: "Blower_drive_on_Q0_7"},
		{ID: "Q1.0", Description: "Collective trouble signal", DataKey: "Collective_trouble_signal_Q1_0"},
		{ID: "Q1.1", Description: "Chiller healthy on", DataKey: "Chiller_healthy_on_Q1_1"},
		{ID: "Q2.0", Description: "Spare", DataKey: "Spare_Q2_0"},
		{ID: "Q2.1", Description: "Condenser fan1 on", DataKey: "Condenser_fan1_on_Q2_1"},
		{ID: "Q2.2", Description: "CR valve 75% on", DataKey: "CR valve 75% on_Q2_2"},
		{ID: "Q2.3", Description: "Chiller fault", DataKey: "Chiller_fault_Q2_3"},
		{ID: "Q2.4", Description: "Condenser fan2 on", DataKey: "Condenser_fan2_on_Q2_4"},
		{ID: "Q2.5", Description: "Condenser fan3 on", DataKey: "Condenser_fan3_on_Q2_5"},
		{ID: "Q2.6", Description: "Condenser fan4 on", DataKey: "Condenser_fan4_on_Q2_6"},
	}
	for _, name := range []string{"gtpl-137-gt-450t-s7-1200", "gtpl-138-gt-450t-s7-1200"} {
		outputsConfigs[name] = thailandOutputs
	}

	// Grain/Paddy AP machines
	apOutputs := []OutputItem{
		{ID: "Q0.0", Description: "Compressor_on", DataKey: "Compressor_on_Q0_0"},
		{ID: "Q0.1", Description: "Compressor_motor_reset", DataKey: "Compressor_motor_reset_Q0_1"},
		{ID: "Q0.2", Description: "CR_valve_25%_on", DataKey: "CR_valve_25_percent_on_Q0_2"},
		{ID: "Q0.3", Description: "CR_valve_50%_on", DataKey: "CR_valve_50_percent_on_Q0_3"},
		{ID: "Q0.4", Description: "Solenoid_valve_on", DataKey: "Solenoid_valve_on_Q0_4"},
		{ID: "Q0.5", Description: "Hot_gas_valve_on", DataKey: "Hot_gas_valve_on_Q0_5"},
		{ID: "Q0.6", Description: "After_heat_valve_on", DataKey: "After_heat_valve_on_Q0_6"},
		{ID: "Q0.7", Description: "Blower_drive_on", DataKey: "Blower_drive_on_Q0_7"},
		{ID: "Q1.0", Description: "Collective_trouble_signal", DataKey: "Collective_trouble_signal_Q1_0"},
		{ID: "Q1.1", Description: "Chiller_healthy_on", DataKey: "Chiller_healthy_on_Q1_1"},
		{ID: "Q2.0", Description: "Spare", DataKey: "Spare_Q2_0"},
		{ID: "Q2.1", Description: "Condenser_fan1_on", DataKey: "Condenser_fan1_on_Q2_1"},
		{ID: "Q2.2", Description: "CR valve 75% on", DataKey: "CR_valve_75_percent_on_Q2_2"},
		{ID: "Q2.3", Description: "Chiller_fault", DataKey: "Chiller_Fault_Q2_3"},
		{ID: "Q2.4", Description: "Condenser_fan2_on", DataKey: "Condenser_fan2_on_Q2_4"},
		{ID: "Q2.5", Description: "CR_valve_100%_on", DataKey: "CR_valve_100_percent_on_Q2_5"},
		{ID: "Q2.6", Description: "Spare", DataKey: "Spare_Q2_6"},
	}
	for _, name := range []string{
		"gtpl-132-300-ap-s7-1200", "gtpl-139-gt-300ap-s7-1200",
		"gtpl-144-gt-300ap-s7-1200", "gtpl-142-gt-450ap-s7-1200",
		"gtpl-123-gt-450ap", "gtpl-143-gt-450ap-s7-1200",
	} {
		outputsConfigs[name] = apOutputs
	}
}

// GetOutputsConfig returns the outputs configuration for a machine
func GetOutputsConfig(name string) ([]OutputItem, bool) {
	cfg, ok := outputsConfigs[strings.ToLower(name)]
	return cfg, ok
}
