package controller

import (

"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

type JobControllerConfiguration struct {
	// ReconcilerSyncLoopPeriod is the amount of time the reconciler sync states loop
	// wait between two reconciler sync.
	// It is set to 15 sec by default.
	// TODO(cph): maybe we can let it grows by multiple in the future
	// and up to 5 minutes to reduce idle loop.
	// e.g. 15s, 30s, 60s, 120s...

	// Enable gang scheduling by kube-batch
	EnableGangScheduling bool
}
