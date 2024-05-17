package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func IsProd() bool {
	return os.Getenv("PRODUCTION") == "true"
}

func IsDev() bool {
	return !IsProd()
}

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func GetOrDefault(key string, def string) string {
	env, ok := os.LookupEnv(key)
	if ok && IsNotBlank(env) {
		return env
	}
	return def
}

func GetOrDefaultInt(key string, def int) int {
	env, ok := os.LookupEnv(key)
	if ok && IsNotBlank(env) {
		i, err := strconv.Atoi(env)
		if err != nil {
			return def
		}
		return i
	}
	return def
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}
