package main

import libvirt "github.com/libvirt/libvirt-go"

func decodeDomainState(status libvirt.DomainState) string {
	switch status {
	case libvirt.DOMAIN_NOSTATE:
		return "nostate"
	case libvirt.DOMAIN_BLOCKED:
		return "blocked"
	case libvirt.DOMAIN_RUNNING:
		return "running"
	case libvirt.DOMAIN_PAUSED:
		return "paused"
	case libvirt.DOMAIN_SHUTDOWN:
		return "shutdown"
	case libvirt.DOMAIN_CRASHED:
		return "crashed"
	case libvirt.DOMAIN_PMSUSPENDED:
		return "pmsuspended"
	case libvirt.DOMAIN_SHUTOFF:
		return "shutoff"
	}
	return "unknown"
}
