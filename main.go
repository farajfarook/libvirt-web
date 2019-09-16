package main

import (
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
)

func main() {
	fmt.Println("connect")
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		panic(err)
	}

	for _, dom := range doms {
		name, _ := dom.GetName()
		fmt.Println(name)
		xml, _ := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
		fmt.Println(xml)
	}

	fmt.Println("close")
	conn.Close()

}
