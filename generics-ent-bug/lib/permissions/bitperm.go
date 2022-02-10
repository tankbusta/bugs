package permissions

type Permission uint64

const (
	NullPermission Permission = 0
	Read           Permission = 1 << iota
	Write
	Execute
)

var PermLookupTbl = map[string]Permission{
	"Read":    Read,
	"Write":   Write,
	"Execute": Execute,
}

// NOTE: Also while this is unused, commenting it out also gets the generator to work
// Even if the privacy statement is there
func GetPermissionsAsList(permissions Permission) []string {
	return GetPermissionsByMask(permissions, PermLookupTbl)
}
