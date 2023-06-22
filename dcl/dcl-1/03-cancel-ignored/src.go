// Issue 198
// Returned cancel function should be called, not avoided.
// Warning should be given.

package testdata

import (
	"context"
	"time"
)

var bg = context.Background()

func _() context.Context {
	var ctx, _ = context.WithDeadline(bg, time.Now().Add(5*time.Second))
	return ctx
}

//<<<<<213, 277>>>>>
