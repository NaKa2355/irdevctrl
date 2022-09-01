package irdev

import "errors"

var (
	ErrDevIO        = errors.New("device io error")
	ErrInvaildInput = errors.New("invaild input data")
)
