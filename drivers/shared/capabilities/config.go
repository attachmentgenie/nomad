package capabilities

import (
	"github.com/syndtr/gocapability/capability"
)

const (
	// HCLSpecLiteral is an equivalent list to NomadDefaults, expressed as a literal
	// HCL string for use in HCL config parsing.
	HCLSpecLiteral = `["AUDIT_WRITE","CHOWN","DAC_OVERRIDE","FOWNER","FSETID","KILL","MKNOD","NET_BIND_SERVICE","SETFCAP","SETGID","SETPCAP","SETUID","SYS_CHROOT"]`
)

// NomadDefaults is the subset of dockerDefaultCaps that Nomad enables by
// default and is used to compute the set of capabilities to add/drop given
// docker driver configuration.
func NomadDefaults() []string {
	return []string{
		"AUDIT_WRITE",
		"CHOWN",
		"DAC_OVERRIDE",
		"FOWNER",
		"FSETID",
		"KILL",
		"MKNOD",
		"NET_BIND_SERVICE",
		"SETFCAP",
		"SETGID",
		"SETPCAP",
		"SETUID",
		"SYS_CHROOT",
	}
}

// DockerDefaults is a list of Linux capabilities enabled by docker by default
// and is used to compute the set of capabilities to add/drop given docker driver
// configuration, as well as Nomad built-in limitations.
//
// https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities
func DockerDefaults() []string {
	return append(NomadDefaults(), "NET_RAW")
}

// Supported returns the set of capabilities supported by the operating system.
func Supported() *Set {
	s := New(nil)

	last := capability.CAP_LAST_CAP

	// workaround for RHEL6 which has no /proc/sys/kernel/cap_last_cap
	if last == capability.Cap(63) {
		last = capability.CAP_BLOCK_SUSPEND
	}

	// accumulate every capability supported by this system
	for _, c := range capability.List() {
		if c > last {
			continue
		}
		s.Add(c.String())
	}

	return s
}
