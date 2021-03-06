package quickfix

import (
	"fmt"
)

// SessionID is a unique identifer of a Session
type SessionID struct {
	BeginString, TargetCompID, SenderCompID, Qualifier string
}

func (s SessionID) String() string {
	if len(s.Qualifier) > 0 {
		return fmt.Sprintf("%s:%s->%s:%s", s.BeginString, s.SenderCompID, s.TargetCompID, s.Qualifier)
	}

	return fmt.Sprintf("%s:%s->%s", s.BeginString, s.SenderCompID, s.TargetCompID)
}
