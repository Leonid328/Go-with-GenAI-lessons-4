package workerpool

import (
	"context"
	"errors"
	"testing"
	"time"
)

// extra test for timeout, as the homework asked the agent to add one
func TestRunPool_TimeoutReturnsDeadline(t *testing.T) {
	jobs := make(chan Job, 1)
	jobs <- Job{
		ID: "late",
		Fetch: func(ctx context.Context) (int, error) {
			select {
			case <-time.After(800 * time.Millisecond):
				return 1, nil
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		},
	}
	close(jobs)

	start := time.Now()
	out := RunPool(jobs, 3, 80*time.Millisecond)

	var got Result
	n := 0
	for r := range out {
		got = r
		n++
	}

	if n != 1 {
		t.Fatalf("got %d results, want 1", n)
	}
	if !errors.Is(got.Err, context.DeadlineExceeded) {
		t.Fatalf("Err = %v, want context.DeadlineExceeded", got.Err)
	}
	if time.Since(start) > 400*time.Millisecond {
		t.Fatalf("took %v, timeout should abort earlier", time.Since(start))
	}
}
