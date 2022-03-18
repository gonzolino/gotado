package gotado

import "context"

// SetSettings updates the mobile device with the given settings
func (md *MobileDevice) SetSettings(ctx context.Context, settings MobileDeviceSettings) error {
	return md.client.put(ctx, apiURL("homes/%d/mobileDevices/%d/settings", md.home.ID, md.ID), settings)
}

// Delete deletes the mobile device
func (md *MobileDevice) Delete(ctx context.Context) error {
	return md.client.delete(ctx, apiURL("homes/%d/mobileDevices/%d", md.home.ID, md.ID))
}
