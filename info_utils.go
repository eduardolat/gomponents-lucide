package lucide

import "fmt"

// Lookup returns the icon information for icon matching "name". The name may be
// provided as either a slug (e.g. copied from the Lucide site), or natural name.
func Lookup(name string) (*IconInfo, error) {
	info, exists := iconsInfoMap[name]
	if !exists {
		return nil, fmt.Errorf("not found")
	}

	return info, nil
}
