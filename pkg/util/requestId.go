package util

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
)

var (
	pid = os.Getpid()
)

func NewRequestID() string {
	return fmt.Sprintf("%d-%s", os.Getpid(), uuid.NewV4().String())
}
