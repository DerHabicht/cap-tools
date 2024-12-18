package pkg

import (
	"fmt"
	"strings"
)

type ContactPriority int

const (
	Primary ContactPriority = iota
	Secondary
	Emergency
)

func ParseContactPriority(s string) (ContactPriority, error) {
	switch strings.ToLower(s) {
	case "primary":
		return Primary, nil
	case "secondary":
		return Secondary, nil
	case "emergency":
		return Emergency, nil
	default:
		return -1, fmt.Errorf("invalid contact priority: %s", s)
	}
}
