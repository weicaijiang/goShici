package gofish

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Url     string
	Methmod string
	Headers *http.Header
	Body    io.Reader
	Client  http.Client
	Handle  Handle
}

func (r *Request) Do() error {
	request, err := http.NewRequest(r.Methmod, r.Url, r.Body)
	if err != nil {
		return err
	}
	request.Header = *r.Headers
	resp, err := r.Client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("err code %d", resp.StatusCode)
	}
	r.Handle.Worker(resp.Body, r.Url)
	defer resp.Body.Close()
	return nil
}

func NewRequest(methmod, Url, userAgent string, handle Handle, body io.Reader) (*Request, error) {
	_, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}
	hdr := http.Header{}
	if userAgent != "" {
		hdr.Add("User-Agent", userAgent)
	} else {
		hdr.Add("User-Agent", UserAgent)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
	return &Request{
		Url:     Url,
		Methmod: methmod,
		Headers: &hdr,
		Body:    body,
		Client:  client,
		Handle:  handle,
	}, nil
}
