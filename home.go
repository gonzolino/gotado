package gotado

import (
	"context"
	"fmt"
)

// GetState returns the state of the home
func (h *Home) GetState(ctx context.Context) (*HomeState, error) {
	homeState := &HomeState{}
	if err := h.client.get(ctx, apiURL("homes/%d/state", h.ID), homeState); err != nil {
		return nil, err
	}
	return homeState, nil
}

// GetZones returns information about the zones in the home
func (h *Home) GetZones(ctx context.Context) ([]*Zone, error) {
	zones := make([]*Zone, 0)
	if err := h.client.get(ctx, apiURL("homes/%d/zones", h.ID), &zones); err != nil {
		return nil, err
	}
	for _, zone := range zones {
		zone.client = h.client
		zone.home = h
	}
	return zones, nil
}

func (h *Home) GetZone(ctx context.Context, name string) (*Zone, error) {
	zones, err := h.GetZones(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to list zones: %w", err)
	}
	for _, zone := range zones {
		if zone.Name == name {
			return zone, nil
		}
	}
	return nil, fmt.Errorf("unknown zone name '%s'", name)
}

// GetWeather returns weather information at the homes location
func (h *Home) GetWeather(ctx context.Context) (*Weather, error) {
	weather := &Weather{}
	if err := h.client.get(ctx, apiURL("homes/%d/weather", h.ID), weather); err != nil {
		return nil, err
	}
	return weather, nil
}

// GetDevices lists all devices in the home
func (h *Home) GetDevices(ctx context.Context) ([]*Device, error) {
	devices := make([]*Device, 0)
	if err := h.client.get(ctx, apiURL("homes/%d/devices", h.ID), &devices); err != nil {
		return nil, err
	}
	for _, device := range devices {
		device.client = h.client
	}
	return devices, nil
}

// GetInstallations lists all installations in the home
func (h *Home) GetInstallations(ctx context.Context) ([]*Installation, error) {
	installations := make([]*Installation, 0)
	if err := h.client.get(ctx, apiURL("homes/%d/installations", h.ID), &installations); err != nil {
		return nil, err
	}
	return installations, nil
}

// GetMobileDevices lists all mobile devices linked to the home
func (h *Home) GetMobileDevices(ctx context.Context) ([]*MobileDevice, error) {
	mobileDevices := make([]*MobileDevice, 0)
	if err := h.client.get(ctx, apiURL("homes/%d/mobileDevices", h.ID), &mobileDevices); err != nil {
		return nil, err
	}
	for _, mobileDevice := range mobileDevices {
		mobileDevice.client = h.client
		mobileDevice.home = h
	}
	return mobileDevices, nil
}

// GetUsers lists all users and their mobile devices linked to the home
func (h *Home) GetUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := h.client.get(ctx, apiURL("homes/%d/users", h.ID), &users); err != nil {
		return nil, err
	}
	for _, user := range users {
		user.client = h.client
	}
	return users, nil
}

// SetPresenceHome sets the geofencing presence to 'at home'.
func (h *Home) SetPresenceHome(ctx context.Context) error {
	presence := PresenceLock{HomePresence: PresenceHome}
	return h.client.put(ctx, apiURL("homes/%d/presenceLock", h.ID), presence)
}

// SetPresenceAway sets the geofencing presence to 'away'.
func (h *Home) SetPresenceAway(ctx context.Context) error {
	presence := PresenceLock{HomePresence: PresenceAway}
	return h.client.put(ctx, apiURL("homes/%d/presenceLock", h.ID), presence)
}

// SetPresenceAuto enables geofencing auto mode.
func (h *Home) SetPresenceAuto(ctx context.Context) error {
	return h.client.delete(ctx, apiURL("homes/%d/presenceLock", h.ID))
}
