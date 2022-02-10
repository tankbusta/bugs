package permissions

import "testing"

func TestGetPermissionsAsList(t *testing.T) {
	perms := Read | Execute

	permissions := GetPermissionsAsList(perms)
	if len(permissions) != 2 {
		t.Fatalf("unexpected number of permissions")
	}
}
