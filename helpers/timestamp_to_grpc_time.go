package helpers

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Convert time.Time to *timestamppb.Timestamp
func TimeToProto(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
