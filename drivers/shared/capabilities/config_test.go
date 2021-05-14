package capabilities

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syndtr/gocapability/capability"
)

// supportedCaps is the legacy implementation of supportedCaps used by the exec
// driver. This is here to ensure the Supported function returns the same
// results.
func legacy_supportedCaps() []string {
	allCaps := []string{}
	last := capability.CAP_LAST_CAP
	// workaround for RHEL6 which has no /proc/sys/kernel/cap_last_cap
	if last == capability.Cap(63) {
		last = capability.CAP_BLOCK_SUSPEND
	}
	for _, cap := range capability.List() {
		if cap > last {
			continue
		}
		allCaps = append(allCaps, fmt.Sprintf("CAP_%s", strings.ToUpper(cap.String())))
	}
	return allCaps
}

func TestConfig_Supported(t *testing.T) {
	s := Supported()
	modern := s.Slice(true)
	fmt.Println("m:", modern)

	legacy := legacy_supportedCaps()
	sort.Strings(legacy)
	fmt.Println("l:", legacy)

	require.Equal(t, legacy, modern)
}
