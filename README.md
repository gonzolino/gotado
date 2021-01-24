# gotado

[![Build Actions Status](https://github.com/gonzolino/gotado/workflows/Build/badge.svg)](https://github.com/gonzolino/gotado/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/gonzolino/gotado)](https://goreportcard.com/report/github.com/gonzolino/gotado) [![PkgGoDev](https://pkg.go.dev/badge/github.com/gonzolino/gotado)](https://pkg.go.dev/github.com/gonzolino/gotado)

Go client for the tado° Web API. Still in early development, so only a subset of the API functionality is implemented.

## Installation

Install gotado with `go get`:

```sh
go get github.com/gonzolino/gotado
```

Then you can import `"github.com/gonzolino/gotado"` in your packages. Have a look at the [examples](examples) directory to see example code.

## Usage

### Authentication

Besides your tado° username and password you need a `clientId` and `clientSecret` to authenticate with the tado° API. One way to get those is to grab them from <https://my.tado.com/webapp/env.js>.

### Code Examples

Get started by creating an authenticated client object:

```golang
client, err := gotado.NewClient("cliendId", "clientSecret").WithCredentials(context.TODO(), "username", "password")
```

With the client you can start using the gotado functions:

```golang
user, err := gotado.GetMe(client)
fmt.Printf("User Email: %s\n", user.Email)

home, err := gotado.GetHome(client, &user.Homes[0])
fmt.Printf("Home Address:\n%s\n%s %s\n", *home.Address.AddressLine1, *home.Address.ZipCode, *home.Address.City)
```

Just have a look at the package documentation to learn more about whats possible.

## Contributing

Please feel free to submit issues and/or pull requests if you discover bugs or missing features.
