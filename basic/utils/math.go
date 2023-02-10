package utils

import (
	"golang.org/x/exp/constraints"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func Add[T constraints.Ordered](x, y T) T {
	return x + y
}
