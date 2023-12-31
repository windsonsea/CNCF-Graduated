package datastore

import (
	"github.com/spiffe/spire/pkg/common/telemetry"
)

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartCountBundleCall return metric
// for server's datastore, on counting bundles.
func StartCountBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Count)
}

// StartAppendBundleCall return metric
// for server's datastore, on sets the bundle.
func StartAppendBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Append)
}

// StartCreateBundleCall return metric
// for server's datastore, on creating a bundle.
func StartCreateBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Create)
}

// StartDeleteBundleCall return metric
// for server's datastore, on deleting a bundle.
func StartDeleteBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Delete)
}

// StartFetchBundleCall return metric
// for server's datastore, on fetching a bundle.
func StartFetchBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Fetch)
}

// StartListBundleCall return metric
// for server's datastore, on listing bundles.
func StartListBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.List)
}

// StartPruneBundleCall return metric
// for server's datastore, on pruning a bundle.
func StartPruneBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Prune)
}

// StartSetBundleCall return metric
// for server's datastore, on sets the bundle.
func StartSetBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Set)
}

// StartUpdateBundleCall return metric
// for server's datastore, on updating a bundle.
func StartUpdateBundleCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.Update)
}

// StartTaintKeyCall return metric
// for server's datastore, on tainting an X.509 CA by key.
func StartTaintX509CAByKeyCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.X509, telemetry.Taint)
}

// StartTaintJWTKeyCall return metric
// for server's datastore, on tainting a JWT public key.
func StartTaintJWTKeyCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.JWT, telemetry.Taint)
}

// StartRevokeX509CACall return metric
// for server's datastore, on revoking an X.509 CA from bundle.
func StartRevokeX509CACall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.X509, telemetry.Revoke)
}

// StartRevokeJWTKeyCall return metric
// for server's datastore, on revoking a JWT Signing Key from bundle.
func StartRevokeJWTKeyCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.Datastore, telemetry.Bundle, telemetry.JWT, telemetry.Revoke)
}

// End Call Counters
