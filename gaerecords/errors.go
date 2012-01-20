package gaerecords

import (
	"os"
)

// The following errors are defined in this package:
var (

	// Error returned when an event callback cancels an operation
	ErrOperationCancelledByEventCallback = os.NewError("gaerecords: Operation was cancelled by an event callback.")
)
