package core

import "context"

// ServiceFunc is a generic function type designed to be used in a handlers as a service function call.
type ServiceFunc[In any, Out any] func(context.Context, In) (Out, error)
