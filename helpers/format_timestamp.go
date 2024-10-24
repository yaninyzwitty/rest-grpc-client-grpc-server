package helpers

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoToTime(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}
