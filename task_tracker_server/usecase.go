package task_tracker_server

import "context"

// Usecase ...
type Usecase interface {
	HelloWorld(ctx context.Context)
}
