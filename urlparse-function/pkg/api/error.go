package api

type RequestError struct {
	StatusCode int
	Body       string
	Err        string
}

func (r RequestError) Error() string {
	return r.Err
}
