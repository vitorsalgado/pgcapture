package source

import (
	"github.com/rueian/pgcapture/pkg/pb"
)

type Source interface {
	Setup() error
	Capture(lsn uint64) (changes chan *pb.Message, err error)
	Commit(lsn uint64)
	Stop()
}
