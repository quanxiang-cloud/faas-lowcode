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

func genFingerprint(ctx context.Context) string {
	once := sync.Once{}
	once.Do(func() {
		key = []byte(strconv.Itoa(rand.Intn(math.MaxInt)))
	})

	val := ctx.Value("User-Id").(string)
	hash := md5.New()
	hash.Write([]byte(val))

	return hex.EncodeToString(hash.Sum(key))
}
