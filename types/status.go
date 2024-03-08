package types

type Status int

const (
	PENDING     = Status(0)
	IN_PROGRESS = Status(1)
	COMPLETED   = Status(2)
	ERROR       = Status(3)
)
