package gotado

import "context"

type Tado struct {
	client *client
}

// New creates a new tado client.
func New(clientID, clientSecret string) *Tado {
	return &Tado{
		client: newClient(clientID, clientSecret),
	}
}

func (t *Tado) authenticate(ctx context.Context, username, password string) error {
	if client, err := t.client.WithCredentials(ctx, username, password); err != nil {
		return err
	} else {
		t.client = client
		return nil
	}
}

// Me authenticates with the given credentials and returns information about the
// authenticated user.
func (t *Tado) Me(ctx context.Context, username, password string) (*User, error) {
	if err := t.authenticate(ctx, username, password); err != nil {
		return nil, err
	}

	me := &User{client: t.client}
	if err := t.client.get(ctx, apiURL("me"), me); err != nil {
		return nil, err
	}
	return me, nil
}
