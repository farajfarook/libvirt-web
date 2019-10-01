# libvirt-web

Simple libvirt (Virtualisation) web API written in golang

### Docker usage

Connect to local libvirt service

`docker run -v /var/run/libvirt/libvirt-sock:/var/run/libvirt/libvirt-sock enbiso/libvirt-web`

Connect to remote libvirt service

`docker run enbiso/libvirt-web --uri 10.10.10.10:5600`

Provide API port/IP

`--addr 192.168.0.0:8081`

### Steps to build

1. Install swag

`go get github.com/swaggo/swag/cmd/swag`

2. Generate swagger docs

`swag init`

3. Install dependencies

`go get -u .`

4. Build project

`go build`