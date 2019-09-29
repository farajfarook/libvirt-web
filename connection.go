package main

import (
	"github.com/libvirt/libvirt-go"
)

//Connection connection wrapper
type Connection struct {
	virt *libvirt.Connect
}

func newConnection(uri string) (*Connection, error) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return nil, err
	}
	return &Connection{
		virt: conn,
	}, nil
}
