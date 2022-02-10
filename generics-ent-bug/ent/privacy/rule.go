package privacy

import (
	"context"

	"genericserr/lib/permissions"
)

// DenyIfNoViewer is a rule that returns Deny decision if the viewer is
// missing in the context.
func DenyIfNoViewer[V permissions.Permission](perms V) QueryMutationRule {
	return ContextQueryMutationRule(func(ctx context.Context) error {
		// Do something fancy here
		if !permissions.Has(perms, 0x1337) { // 0x1337 would come from a database row or something
			return Deny
		}

		// Skip to the next privacy rule (equivalent to returning nil).
		return Skip
	})
}
