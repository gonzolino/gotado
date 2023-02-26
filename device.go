package gotado

import (
	"context"
	"fmt"
)

// GetTemperatureOffset returns the temperature offsets of the device
func (d *Device) GetTemperatureOffset(ctx context.Context) (float32, error) {
	temperatureOffset := &TemperatureOffset{}
	if err := d.client.get(ctx, apiURL("devices/%s/temperatureOffset", d.SerialNo), &temperatureOffset); err != nil {
		return 0, err
	}
	switch d.home.TemperatureUnit {
	case TemperatureUnitCelsius:
		return temperatureOffset.Celsius, nil
	case TemperatureUnitFahrenheit:
		return temperatureOffset.Fahrenheit, nil
	default:
		return 0, fmt.Errorf("unknown temperature unit: %s", d.home.TemperatureUnit)
	}
}

// SetTemperatureOffset updates the temperature offsets of the device
func (d *Device) SetTemperatureOffset(ctx context.Context, temperatureOffset float32) error {
	offset := &TemperatureOffset{}
	switch d.home.TemperatureUnit {
	case TemperatureUnitCelsius:
		offset.Celsius = temperatureOffset
	case TemperatureUnitFahrenheit:
		offset.Fahrenheit = temperatureOffset
	}
	return d.client.put(ctx, apiURL("devices/%s/temperatureOffset", d.SerialNo), offset)
}
