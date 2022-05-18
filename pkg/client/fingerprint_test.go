package client

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestFingerprint(t *testing.T) {
	if !through() {
		t.Failed()
	}
}

func through() bool {
	ctx := context.Background()
	var userID interface{} = "User-Id"
	ctx = context.WithValue(ctx, userID, uuid.New().String())

	ctx = WithFingerprint(ctx)
	return isAllowedThrough(ctx)
}

func TestScramble(t *testing.T) {
	ctx := context.Background()
	var userID interface{} = "User-Id"
	ctx = context.WithValue(ctx, userID, uuid.New().String())
	ctx = WithFingerprint(ctx)
	ctx = context.WithValue(ctx, userID, uuid.New().String())
	if isAllowedThrough(ctx) {
		t.Failed()
	}
}

func BenchmarkFingerprint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !through() {
			b.Failed()
		}
	}
}
