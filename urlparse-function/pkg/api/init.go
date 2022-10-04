package api

import "net/http"

func New(options Options) api {
	return api{
		Options: options,
		Client: http.Client{
			Transport: &MyJWTClient{
				password:  options.Password,
				loginUrl:  options.LoginUrl,
				transport: http.DefaultTransport,
			},
		},
	}
}
