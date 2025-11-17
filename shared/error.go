package shared

type Error struct {
	Code    int
	Message string
	Details string
}

func (e *Error) Error() string {
	return e.Message
}
