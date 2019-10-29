package main

import (
	nut "github.com/robbiet480/go.nut"
)

const serialField = "ups.serial"
const percentField = "ups.load"
const wattsField = "ups.realpower.nominal"

var valueFields = []string{
	"battery.charge",
	"battery.runtime",
	"battery.voltage",

	"input.voltage",

	"output.voltage",

	percentField,
	wattsField,
}

type mappedUPS struct {
	ups         *nut.UPS
	VariableMap map[string]nut.Variable
}

func mapUPS(ups *nut.UPS) mappedUPS {
	variableMap := map[string]nut.Variable{}

	for _, variable := range ups.Variables {
		variableMap[variable.Name] = variable
	}

	return mappedUPS{
		ups:         ups,
		VariableMap: variableMap,
	}
}

type influxableUPS mappedUPS

func (device influxableUPS) Tags() map[string]string {
	tags := map[string]string{}

	if variable, ok := device.VariableMap[serialField]; ok {
		tags[variable.Name] = variable.Value.(string)
	}

	return tags
}

func (device influxableUPS) Fields() map[string]interface{} {
	fields := map[string]interface{}{}

	for _, valueField := range valueFields {
		if variable, ok := device.VariableMap[valueField]; ok {
			fields[variable.Name] = variable.Value
		}
	}

	percent, hasPercent := fields[percentField]
	watts, hasWatts := fields[wattsField]

	if hasPercent && hasWatts {
		percentFloat := percent.(int64)
		wattsFloat := watts.(int64)

		wattUsage := (wattsFloat * percentFloat) / 100
		fields["estimated-watts"] = wattUsage
	}

	return fields
}

func (device influxableUPS) Category() string {
	return "ups"
}
