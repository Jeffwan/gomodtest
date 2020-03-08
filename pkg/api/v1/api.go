package v1

import (
   "k8s.io/client-go/tools/cache"

)


var (
	// KeyFunc is the short name to DeletionHandlingMetaNamespaceKeyFunc.
	// IndexerInformer uses a delta queue, therefore for deletes we have to use this
	// key function but it should be just fine for non delete events.
	KeyFunc = cache.DeletionHandlingMetaNamespaceKeyFunc

)

type ReplicaStatus struct {
	// The number of actively running pods.
	Active int32

	// The number of pods which reached phase Succeeded.
	Succeeded int32

	// The number of pods which reached phase Failed.
	Failed int32
}
