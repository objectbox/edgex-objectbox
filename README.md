# EdgeX ObjectBox Edition

[![Go Report Card](https://goreportcard.com/badge/github.com/objectbox/edgex-objectbox)](https://goreportcard.com/report/github.com/objectbox/edgex-objectbox)
[![license](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)

[EdgeX](https://www.edgexfoundry.org/) is the open source edge computing platform, and this repo contains a EdgeX version running on the [ObjectBox](https://objectbox.io/) database.
ObjectBox is a resource-efficient and fast embedded database and thus is particular well suited for constrained edge devices.
This repo is provided and maintained by ObjectBox. 

*Original text for EdgeX:*

EdgeX Foundry is a vendor-neutral open source project hosted by The Linux Foundation building a common open framework for IoT edge computing.  At the heart of the project is an interoperability framework hosted within a full hardware- and OS-agnostic reference software platform to enable an ecosystem of plug-and-play components that unifies the marketplace and accelerates the deployment of IoT solutions.  This repository contains the Go implementation of EdgeX Foundry microservices.  It also includes files for building the services, containerizing the services, and initializing (bootstrapping) the services.

# Get Started

ObjectBox provides docker images in our organization's [DockerHub repository](https://hub.docker.com/u/objectboxio/).
They can be launched easily with **docker-compose**.

The simplest way to get started is to fetch the latest docker-compose.yml and start the EdgeX containers:

```sh
wget https://raw.githubusercontent.com/objectbox/edgex-objectbox/fuji/bin/docker-compose.yml
docker-compose up -d
```

You can check the status of your running EdgeX services by going to http://localhost:8500/

Now that you have EdgeX up and running, you can follow our [API Walkthrough](https://docs.edgexfoundry.org/Ch-Walkthrough.html) to learn how the different services work together to connect IoT devices to cloud services.

### Without security services
The following docker-compose file omits security services for an even more lightweight setup.  
Note: this is the only option available for armhf/arm-32bit devices/OSs, such as Raspberry Pi with Raspbian, 
because Kong API gateway is not available for armhf.

```sh
wget https://raw.githubusercontent.com/objectbox/edgex-objectbox/fuji/bin/docker-compose-no-secty.yml
docker-compose -f docker-compose-no-secty.yml up -d
```

# Running EdgeX with security components

Starting with the Fuji release, EdgeX includes enhanced security features that are enabled by default. There are 3 major components that are responsible for security
features: 

- Security-secrets-setup
- Security-secretstore-setup
- Security-proxy-setup

When security features are enabled, additional steps are required to access the resources of EdgeX.

1. The user needs to create an access token and associate every REST request with the access token. 
2. The exported external ports (such as 48080, 48081 etc.) will be inaccessible for security purposes. Instead, all REST requests need to go through the proxy. The proxy will redirect the request to the individual microservices on behalf of the user.

Sample steps to create an access token and use the token to access EdgeX resources can be found here: [Security Components](SECURITY.md)

# Other installation and deployment options

## Snap Package

EdgeX Foundry is also available as a snap package, for more details
on the snap, including how to install it, please refer to [EdgeX snap](snap/README.md)

## Native binaries

### Prerequisites

#### Go

- The current targeted version of the Go language runtime for release artifacts is v1.12.x
- The minimum supported version of the Go language runtime is v1.11.5

#### ObjectBox

To install the ObjectBox dynamic library (.so, .dylib, .dll), run this in your terminal (Windows users, please use something like Git Bash):
 
```bash
bash <(curl -s https://raw.githubusercontent.com/objectbox/objectbox-c/master/download.sh)

```

Alternatively, the [ObjectBox C repository](https://github.com/objectbox/objectbox-c) comes with additional options and details on this procedure. 

#### pkg-config

`go get github.com/rjeczalik/pkgconfig/cmd/pkg-config`

#### ZeroMQ

Several EdgeX Foundry services depend on ZeroMQ for communications by default.

The easiest way to get and install ZeroMQ on Linux is to use this [setup script](https://gist.github.com/katopz/8b766a5cb0ca96c816658e9407e83d00).

For macOS, use brew:

```sh
brew install zeromq
```

For directions installing ZeroMQ on Windows, please see [the Windows documentation.](ZMQWindows.md)

#### pkg-config

The necessary file will need to be added to the `PKG_CONFIG_PATH` environment variable.

On Linux, add this line to your local profile:

```sh
export PKG_CONFIG_PATH=/usr/local/Cellar/zeromq/4.2.5/lib/pkgconfig/
```

For macOS, install the package with brew:

```sh
brew install pkg-config
```

### Installation and Execution

EdgeX is organized as Go Modules; there is no requirement to set the GOPATH or
GO111MODULE envrionment variables nor is there a requirement to root all the components under ~/go
(or $GOPATH) and use the `go get` command. In other words,

```sh
git clone git@github.com/objectbox/edgex-objectbox.git
cd edgex-go
make build
```

If you do want to root everthing under $GOPATH, you're free to use that pattern as well

```sh
GO111MODULE=on && export GO111MODULE
go get github.com/objectbox/edgex-objectbox
cd $GOPATH/src/github.com/objectbox/edgex-objectbox
make build
```

To start EdgeX

```sh
make run
```

or

```sh
cd bin
./edgex-launch.sh
```

**Note** You must have a database (Mongo or Redis) running before the services will operate
correctly. If you don't want to install a database locally, you can host one via Docker. You may
also need to change the `configuration.toml` files for one or more of the services.

## Build your own Docker Containers

In addition to running the services directly, Docker and Docker Compose can be used.

### Prerequisites

See [the install instructions](https://docs.docker.com/install/) to learn how to obtain and install Docker.

### Installation and Execution

Follow the "Installation and Execution" steps above for obtaining and building the code, then

```sh
make docker run_docker
```


# Community

- Chat: [https://edgexfoundry.slack.com](https://join.slack.com/t/edgexfoundry/shared_invite/enQtNDgyODM5ODUyODY0LWVhY2VmOTcyOWY2NjZhOWJjOGI1YzQ2NzYzZmIxYzAzN2IzYzY0NTVmMWZhZjNkMjVmODNiZGZmYTkzZDE3MTA)
- Mailing lists: https://lists.edgexfoundry.org/mailman/listinfo

# License

[Apache-2.0](LICENSE)
