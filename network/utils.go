package domain

import libvirt "github.com/libvirt/libvirt-go"

var conn *libvirt.Connect

//Init network
func Init(connect *libvirt.Connect) {
	conn = connect
}
