package gotado

import (
	"context"
	"fmt"
)

// GetHome returns the home with the given name
func (u *User) GetHome(ctx context.Context, name string) (*Home, error) {
	var homeID int32
	for _, h := range u.Homes {
		if h.Name == name {
			homeID = h.ID
		}
	}
	if homeID == 0 {
		return nil, fmt.Errorf("unknown home name '%s'", name)
	}

	home := &Home{client: u.client}
	if err := u.client.get(ctx, apiURL("homes/%d", homeID), home); err != nil {
		return nil, err
	}
	return home, nil
}
