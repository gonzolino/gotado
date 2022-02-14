package gotado

import "context"

// GetTemperatureOffset returns the temperature offsets of the device
func (d *Device) GetTemperatureOffset(ctx context.Context) (*TemperatureOffset, error) {
	temperatureOffset := &TemperatureOffset{}
	if err := d.client.get(ctx, apiURL("devices/%s/temperatureOffset", d.SerialNo), &temperatureOffset); err != nil {
		return nil, err
	}
	return temperatureOffset, nil
}
