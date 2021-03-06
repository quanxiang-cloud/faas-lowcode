package client

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/rand"
	"strconv"
	"sync"
)

var key []byte

var fingerprintKey interface{} = "Quanxiang-FingerprintKey"

func isAllowedThrough(ctx context.Context) bool {
	key, ok := ctx.Value(fingerprintKey).(string)
	if !ok {
		return false
	}

	return key == genFingerprint(ctx)
}

func WithFingerprint(ctx context.Context) context.Context {
	return context.WithValue(ctx, fingerprintKey, genFingerprint(ctx))
}

var once = sync.Once{}

func genFingerprint(ctx context.Context) string {
	once.Do(func() {
		key = []byte(strconv.Itoa(rand.Intn(math.MaxInt32)))
	})

	var userID string
	val := ctx.Value("User-Id")
	if val != nil {
		userID = ctx.Value("User-Id").(string)
	}

	hash := md5.New()
	hash.Write([]byte(userID))

	return hex.EncodeToString(hash.Sum(key))
}
