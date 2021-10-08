package util

import (
	"github.com/go-basic/uuid"
	"strings"
)

func NewUUID() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
