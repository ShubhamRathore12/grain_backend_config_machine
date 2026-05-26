package config

import "strings"

// allMachines holds the complete configuration for every machine
var allMachines = []Machine{
	// ─── S7-200 Machines ─────────────────────────────────────────────────────
	{
		Name: "GTPL-108-gT-40E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-40E-P", Table: "GTPL_108_gT_40E_P_S7_200_Germany",
		StatusKey: "GTPL_108", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-109-gT-40E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-40E-P", Table: "GTPL_109_gT_40E_P_S7_200_Germany",
		StatusKey: "GTPL_109", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-110-gT-40E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-40E-P", Table: "GTPL_110_gT_40E_P_S7_200_Germany",
		StatusKey: "GTPL_110", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-111-gT-80E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-80E-P", Table: "GTPL_111_gT_80E_P_S7_200_Germany",
		StatusKey: "GTPL_111", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-112-gT-80E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-80E-P", Table: "GTPL_112_gT_80E_P_S7_200_Germany",
		StatusKey: "GTPL_112", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-113-gT-80E-P-S7-200", Location: "Germany", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-80E-P", Table: "GTPL_113_gT_80E_P_S7_200_Germany",
		StatusKey: "GTPL_113", Tags: S7_200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: true,
	},
	{
		Name: "GTPL-118-gT-60T-S7-200", Location: "Telangana", Image: "/images/200.jpg",
		PLC: "S7-200", ChillerModel: "gT-80E-P", Table: "GTPL_118_GT_60T_S7_1200",
		StatusKey: "KABO_200", Tags: GTPL_118_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},

	// ─── S7-1200 Machines (with heater / analog) ─────────────────────────────
	{
		Name: "GTPL-30-gT-180E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-140E", Table: "GTPL_114_GT_140E_S7_1200",
		StatusKey: "GTPL_114", Tags: GPL_115_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},
	{
		Name: "GTPL-115-gT-180E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-180E", Table: "GTPL_115_GT_180E_S7_1200",
		StatusKey: "GTPL_115", Tags: GPL_115_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},
	{
		Name: "GTPL-116-gT-240E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_116_GT_240E_S7_1200",
		StatusKey: "GTPL_116", Tags: S7_1200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},
	{
		Name: "GTPL-117-gT-320E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-320E", Table: "GTPL_117_GT_320E_S7_1200",
		StatusKey: "GTPL_117", Tags: GPL_117_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},
	{
		Name: "GTPL-119-gT-180E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-180E", Table: "GTPL_119_GT_180E_S7_1200",
		StatusKey: "GTPL_119", Tags: S7_1200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},
	{
		Name: "GTPL-120-gT-180E-S7-1200", Location: "Germany", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-180E", Table: "GTPL_120_GT_180E_S7_1200",
		StatusKey: "GTPL_120", Tags: S7_1200_TAGS, MenuType: "standard", AerationType: "withHeating",
		PressureUnit: "psi", CondenserFans: 1, HasHTR: true, HasAnalogMenu: true, IsS7200: false,
	},

	// ─── S7-1200 Large Machines (no heater, 6 fans) ──────────────────────────
	{
		Name: "GTPL-121-gT-1000T-S7-1200", Location: "kanpur", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-1000T", Table: "GTPL_121_GT1000T",
		StatusKey: "GTPL_121", Tags: S7_1200_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-122-gT-1000T-S7-1200", Location: "kanpur", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-1000T", Table: "gtpl_122_s7_1200_01",
		StatusKey: "GTPL_122", Tags: S7_1200_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},

	// ─── GT-650T Machines (6 fans, no heater) ────────────────────────────────
	{
		Name: "GTPL-081-GT-650T-S7-1200", Location: "Dharuhera", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "GT-650T", Table: "GTPL_081_GT_650T_S7_1200",
		StatusKey: "GTPL_081", Tags: GPL_124_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-105-GT-650T-S7-1200", Location: "Dharuhera", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "GT-650T", Table: "GTPL_105_GT_650T_S7_1200",
		StatusKey: "GTPL_105", Tags: GPL_124_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-131-GT-650T-S7-1200", Location: "Ganganagar, Rajasthan", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_131_GT_650T_S7_1200",
		StatusKey: "GTPL_131", Tags: GPL_124_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-133-GT-650T-S7-1200", Location: "Vietnam", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_133_GT_650T_S7_1200",
		StatusKey: "GTPL_133", Tags: GPL_124_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 6, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},

	// ─── GT-450T Machines (4 fans) ───────────────────────────────────────────
	{
		Name: "GTPL-124-GT-450T-S7-1200", Location: "Indonesia", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_124_GT_450T_S7_1200",
		StatusKey: "GTPL_124", Tags: GPL_124_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-134-gT-450T-S7-1200", Location: "Kakinada (AP)", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450T", Table: "GTPL_134_GT_450T_S7_1200",
		StatusKey: "GTPL_134", Tags: GTPL_134_135_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-135-gT-450T-S7-1200", Location: "Bihar", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450T", Table: "GTPL_135_GT_450T_S7_1200",
		StatusKey: "GTPL_135", Tags: GTPL_134_135_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-137-GT-450T-S7-1200", Location: "Thailand", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_137_GT_450T_S7_1200",
		StatusKey: "GTPL_137", Tags: GTPL_137_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "bar", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-138-GT-450T-S7-1200", Location: "Thailand", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_138_GT_450T_S7_1200",
		StatusKey: "GTPL_138", Tags: GTPL_138_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "bar", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-145-gT-450T-S7-1200", Location: "Tamil Nadu", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450T", Table: "GTPL_145_GT_450T_S7_1200",
		StatusKey: "GTPL_145", Tags: GTPL_134_135_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-148-gT-450T-S7-1200", Location: "Tamil Nadu", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450T", Table: "GTPL_148_GT_450T_S7_1200",
		StatusKey: "GTPL_148", Tags: GTPL_134_135_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-061-gT-450T-S7-1200", Location: "Turkey", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450T", Table: "GTPL_061_GT_450T_S7_1200",
		StatusKey: "GTPL_061", Tags: GTPL_134_135_TAGS, MenuType: "standard", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},

	// ─── Grain/Paddy (AP) Machines ───────────────────────────────────────────
	{
		Name: "GTPL-123-GT-450AP", Location: "Raichur, Karnataka", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gt-450AP", Table: "GTPL_123_GT_450AP_S7_1200",
		StatusKey: "GTPL_123", Tags: GPL_132_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-132-300-AP-S7-1200", Location: "Salem (Tamil Nadu)", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-240E", Table: "GTPL_132_GT300AP",
		StatusKey: "GTPL_132", Tags: GPL_132_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-136-gT-450AP", Location: "Srilanka", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450AP", Table: "GTPL_136_GT_450AP_S7_1200",
		StatusKey: "GTPL_136", Tags: GPL_132_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 4, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-139-GT-300AP-S7-1200", Location: "Pondicherry", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "GT-300AP", Table: "GTPL_139_GT300AP",
		StatusKey: "GTPL_139", Tags: GTPL_139_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-142-gT-450AP-S7-1200", Location: "A.P.", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450AP", Table: "GTPL_142_GT_450AP_S7_1200",
		StatusKey: "GTPL_142", Tags: GPL_132_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-143-gT-450AP-S7-1200", Location: "A.P.", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "gT-450AP", Table: "GTPL_143_GT_450AP_S7_1200",
		StatusKey: "GTPL_143", Tags: GPL_132_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
	{
		Name: "GTPL-144-GT-300AP-S7-1200", Location: "Tamil Nadu", Image: "/images/1200.jpg",
		PLC: "S7-1200", ChillerModel: "GT-300AP", Table: "GTPL_144_GT_300AP_S7_1200",
		StatusKey: "GTPL_144", Tags: GTPL_139_TAGS, MenuType: "grainPaddy", AerationType: "withoutHeating",
		PressureUnit: "psi", CondenserFans: 2, HasHTR: false, HasAnalogMenu: false, IsS7200: false,
	},
}

// machineIndex is a lookup map built at init time
var machineIndex map[string]*Machine

func init() {
	machineIndex = make(map[string]*Machine, len(allMachines))
	for i := range allMachines {
		machineIndex[strings.ToLower(allMachines[i].Name)] = &allMachines[i]
	}
}

// GetAllMachines returns the full list of machines
func GetAllMachines() []Machine {
	return allMachines
}

// GetMachineByName looks up a machine by name (case-insensitive)
func GetMachineByName(name string) (Machine, bool) {
	m, ok := machineIndex[strings.ToLower(name)]
	if !ok {
		return Machine{}, false
	}
	return *m, true
}

// GetMachineTags returns the fault tags for a machine
func GetMachineTags(name string) ([]string, bool) {
	m, ok := machineIndex[strings.ToLower(name)]
	if !ok {
		return nil, false
	}
	return m.Tags, true
}
