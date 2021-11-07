package log

import (
	"github.com/hashicorp/raft"
)

type Config struct {
	Raft struct {
		raft.Config
		StreamLayer *StreamLayer
		// bootstrap a server configured with itself as the only voter
		// wait until it becomes the leader, and then tell the leader to add more servers to the cluster
		// the subsequently added servers don't bootstrap
		Bootstrap bool
	}
	Segment struct {
		MaxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
	}
}
