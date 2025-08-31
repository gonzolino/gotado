# gotado

[![CodeQL](https://github.com/gonzolino/gotado/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/gonzolino/gotado/actions/workflows/github-code-scanning/codeql) [![Go Report Card](https://goreportcard.com/badge/github.com/gonzolino/gotado)](https://goreportcard.com/report/github.com/gonzolino/gotado) [![PkgGoDev](https://pkg.go.dev/badge/github.com/gonzolino/gotado)](https://pkg.go.dev/github.com/gonzolino/gotado)

Go client for the tado° Web API. Still in early development, so only a subset of the API functionality is implemented.

## Installation

Install gotado with `go get`:

```sh
go get github.com/gonzolino/gotado
```

Then you can import `"github.com/gonzolino/gotado"` in your packages. Have a look at the [examples](examples) directory to see example code.

## Usage

### Authentication

Authentication to the tado° API is done using OAuth 2.0. The details are documented [here](https://support.tado.com/en/articles/8565472-how-do-i-authenticate-to-access-the-rest-api).

Gotado requires an `oauth2.Config` and a `oauth2.Token` to make authenticated requests. An authentication flow on the CLI could look like this:

```golang
ctx := context.Background()
config := gotado.AuthConfig(clientID)

deviceAuth, err := config.DeviceAuth(ctx)

// The user must visit the verification URI and authenticate with his personal credentials.
// Afterwards an access token will be generated.
fmt.Printf("To authenticate, visit %s\n", deviceAuth.VerificationURIComplete)

token, err := config.DeviceAccessToken(ctx, deviceAuth)
```

Be aware that an access token is only valid for 10 minutes. To keep authenticated for longer, you will need to request a refresh token during authentication. You can do this by adding the `offline_access` scope to the auth config:

```golang
config := gotado.AuthConfig(clientID, "offline_access")
```

Gotado will automatically refresh the access token when it expires, if a refresh token is available.

### Getting Started

Get started by creating a client object:

```golang
tado := gotado.New(ctx, config, token)
```

With the client you can start using the gotado functions:

```golang
me, err := tado.Me(ctx)
fmt.Printf("User Email: %s\n", me.Email)

home, err := me.GetHome(ctx, "My Home Name")
fmt.Printf("Home Address:\n%s\n%s %s\n", *home.Address.AddressLine1, *home.Address.ZipCode, *home.Address.City)
```

Just have a look at the package documentation to learn more about whats possible.

## Contributing

Please feel free to submit issues and/or pull requests if you discover bugs or missing features.
