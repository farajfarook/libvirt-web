package main

import (
	"fmt"
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

func main() {
	fmt.Println("Hello")
	var drive uint
	drive = 0
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("failed to connect to qemu")
	}
	defer conn.Close()
}
