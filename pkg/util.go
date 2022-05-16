package pkg

import (
	"os"
	"strconv"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvWithDefault(key, s string) string {
	env := GetEnv(key)
	if env == "" {
		env = s
	}

	return env
}

func GetEnvToInt64(key string) int64 {
	env := os.Getenv(key)
	i, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		i = 0
	}

	return i
}

func GetEnvToInt64WithDefatult(key string, si int64) int64 {
	i := GetEnvToInt64(key)
	if i == 0 {
		i = si
	}

	return i
}
