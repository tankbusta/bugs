package permissions

type (
	PermissionT interface {
		Permission | uint64
	}
)

// Has checks if a bit exists on the mask
func Has[V PermissionT](b, flag V) bool { return b&flag != 0 }

func GetPermissionsByMask[V PermissionT](permissions V, table map[string]V) []string {
	if permissions == 0 {
		return nil
	}

	out := make([]string, 0)
	for capability, bit := range table {
		if Has(permissions, bit) {
			out = append(out, capability)
		}
	}

	return out
}
