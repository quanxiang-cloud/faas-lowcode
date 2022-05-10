package client

import (
	"os"
	"strconv"
)

func getEnv(key string) string {
	return os.Getenv(key)
}

func getEnvWithDefault(key, s string) string {
	env := getEnv(key)
	if env == "" {
		env = s
	}

	return env
}

func getEnvToInt64(key string) int64 {
	env := os.Getenv(key)
	i, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		i = 0
	}

	return i
}

func getEnvToInt64WithDefatult(key string, si int64) int64 {
	i := getEnvToInt64(key)
	if i == 0 {
		i = si
	}

	return i
}
