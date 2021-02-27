package errory

func ErrorHandler(err error) (int, interface{}) {
	if err == nil {
		return unknownErrorCode, "unknown error"
	}

	if e, ok := err.(*Error); ok {
		return e.code, e.msg
	} else {
		return defaultErrorCode, err.Error()
	}
}
