package client

import "context"

// predefined header
const (
	RequestID = "Request-Id"

	Timezone = "Timezone"

	TenantID = "Tenant-Id"

	userID = "User-Id"
)

type KV []string

func (k KV) Wreck() (string, string) {
	switch len(k) {
	case 0:
		return "", ""
	case 1:
		return k[0], ""
	default:
		return k[0], k[1]
	}
}

// Fuzzy return key and value as []interface
func (k KV) Fuzzy() (result []interface{}) {
	for _, elem := range k {
		result = append(result, elem)
	}
	return
}

// GetRequestIDKV return request id
func GetRequestIDKV(ctx context.Context) KV {
	i := ctx.Value(RequestID)
	rid, ok := i.(string)
	if ok {
		return KV{RequestID, rid}
	}
	return KV{RequestID, "unexpected type"}
}

// GetTimezone return timezone
func GetTimezone(ctx context.Context) KV {
	i := ctx.Value(Timezone)
	tz, ok := i.(string)
	if ok {
		return KV{Timezone, tz}
	}
	return KV{Timezone, "unexpected type"}
}

// GetTenantID return tenantID
func GetTenantID(ctx context.Context) KV {
	i := ctx.Value(TenantID)
	tid, ok := i.(string)
	if ok {
		return KV{TenantID, tid}
	}
	return KV{TenantID, "unexpected type"}
}

// GetUserID return userID
func GetUserID(ctx context.Context) KV {
	i := ctx.Value(userID)
	tid, ok := i.(string)
	if ok {
		return KV{userID, tid}
	}
	return KV{userID, "unexpected type"}
}
