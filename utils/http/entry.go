package http

import "io"

type Client interface {
	SetHeader(k, v string)
	Get(uri string, args map[string]string)
	Post(uri string, params io.Reader)
}
