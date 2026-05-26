package config

import "strings"

// GetMenuConfig returns the menu configuration for a machine
func GetMenuConfig(name string) (*MenuConfig, bool) {
	m, ok := machineIndex[strings.ToLower(name)]
	if !ok {
		return nil, false
	}

	menu := &MenuConfig{
		MenuType:   m.MenuType,
		ShowAnalog: m.HasAnalogMenu,
	}

	if m.MenuType == "grainPaddy" {
		menu.Items = []MenuItem{
			{Title: "Grain Chilling Mode", Path: "auto-grain", Icon: "Factory", Gradient: "from-emerald-500 via-teal-500 to-green-600"},
			{Title: "Paddy Ageing Mode", Path: "auto-paddy", Icon: "Gauge", Gradient: "from-amber-500 via-yellow-500 to-orange-600"},
			{Title: "Aeration", Path: "aerations", Icon: "Wind", Gradient: "from-sky-500 via-cyan-500 to-blue-600"},
			{Title: "Fault", Path: "fault", Icon: "AlertTriangle", Gradient: "from-red-500 via-orange-500 to-rose-600"},
			{Title: "Settings", Path: "settings", Icon: "Cog", Gradient: "from-slate-500 via-zinc-500 to-gray-600"},
			{Title: "Inputs", Path: "inputs", Icon: "Power", Gradient: "from-violet-500 via-indigo-500 to-purple-600"},
			{Title: "Outputs", Path: "outputs", Icon: "Zap", Gradient: "from-fuchsia-500 via-pink-500 to-rose-600"},
			{Title: "Analog", Path: "inputs/analog", Icon: "Activity", Gradient: "from-teal-500 via-emerald-500 to-green-600"},
			{Title: "Test", Path: "test", Icon: "Construction", Gradient: "from-cyan-500 via-blue-500 to-indigo-600"},
		}
	} else {
		menu.Items = []MenuItem{
			{Title: "Auto", Path: "auto", Icon: "Gauge", Gradient: "from-indigo-500 via-blue-500 to-purple-600"},
			{Title: "Aeration", Path: "aerations", Icon: "Wind", Gradient: "from-sky-500 via-cyan-500 to-blue-600"},
			{Title: "Fault", Path: "fault", Icon: "AlertTriangle", Gradient: "from-red-500 via-orange-500 to-rose-600"},
			{Title: "Settings", Path: "settings", Icon: "Cog", Gradient: "from-slate-500 via-zinc-500 to-gray-600"},
			{Title: "Inputs", Path: "inputs", Icon: "Power", Gradient: "from-violet-500 via-indigo-500 to-purple-600"},
			{Title: "Outputs", Path: "outputs", Icon: "Zap", Gradient: "from-fuchsia-500 via-pink-500 to-rose-600"},
			{Title: "Test", Path: "test", Icon: "Construction", Gradient: "from-cyan-500 via-blue-500 to-indigo-600"},
		}

		// Add analog menu for machines that support it
		if m.HasAnalogMenu {
			menu.Items = append(menu.Items, MenuItem{
				Title: "Analog", Path: "inputs/analog", Icon: "Activity", Gradient: "from-teal-500 via-emerald-500 to-green-600",
			})
		}
	}

	return menu, true
}
