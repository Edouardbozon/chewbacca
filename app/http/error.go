package http

type internalErr struct {
	msg string
}

func (e *internalErr) Error() string {
	return e.msg
}
