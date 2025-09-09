package types

import (
	"os"
)

type EnvVar string

const (
	ServerPort EnvVar = "SERVER_PORT"
	JWTSecret  EnvVar = "JWT_SECRET"
)

func (e EnvVar) Value() string {
	return os.Getenv(string(e))
}

func (e EnvVar) String() string {
	return string(e)
}
