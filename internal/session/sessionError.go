package session

type SessionHasNoData struct {
}

func (e *SessionHasNoData) Error() string {
	return "Session had no data after get"
}
