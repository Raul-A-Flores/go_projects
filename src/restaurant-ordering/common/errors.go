package common

import "errors"

var (
	ErrorNoItems = errors.New("items must have at least one item")
)
