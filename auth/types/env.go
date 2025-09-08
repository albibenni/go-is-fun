package types

import (
	"os"
)

type EnvVar string

const (
	ServerPath EnvVar = "SERVER_PATH"
)

func (e EnvVar) Value() string {
	return os.Getenv(string(e))
}

func (e EnvVar) String() string {
	return string(e)
}
