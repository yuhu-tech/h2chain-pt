# Createorder Service

This is the Createorder service

Generated with

```
micro new H2chainProject/createorder --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.createorder
- Type: api
- Alias: createorder

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./createorder-api
```

Build a docker image
```
make docker
```