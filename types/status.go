package types

type Status int

const (
	PENDING     = Status(0)
	IN_PROGRESS = Status(1)
	COMPLETED   = Status(2)
	ERROR       = Status(3)
)

func (s Status) IsPending() bool {
	return s == PENDING
}

func (s Status) HasCompleted() bool {
	return s == COMPLETED
}

func (s Status) HasError() bool {
	return s == ERROR
}
