//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"context"
	"sync"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	rw        sync.RWMutex // unnecessary for this task, but good practice
	TimeUsed  time.Duration
}

const freemiumLimit = time.Second * 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(context.Context), u *User) bool {
	ctx := context.Background()
	var cancel context.CancelFunc

	u.rw.RLock()
	remaining := freemiumLimit - u.TimeUsed
	u.rw.RUnlock()
	// log.Println("REMAINING TIME: ", remaining)

	done := make(chan struct{})

	if !u.IsPremium {
		ctx, cancel = context.WithTimeout(ctx, remaining)
		defer cancel()
	}

	go func(c context.Context) {
		// process() // to kill the process this function should be context friendly
		process(c)
		close(done)
	}(ctx)

	select {
	case <-ctx.Done():
		return false
	case <-done:
		return true
	}
}

func main() {
	RunMockServer()
}
