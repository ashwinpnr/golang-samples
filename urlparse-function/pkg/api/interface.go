package api

type APIInterface interface {
	DoGetRespose(requestURL string) (Response, error)
}

type Response interface {
	GetResponse() string
}
