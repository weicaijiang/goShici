package gofish

import "time"

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 Edg/87.0.664.75"
	Qps       = 50
)

var rateLimiter = time.Tick(time.Second / Qps)

type GoFish struct {
	Request *Request
}

func NewGoFish() *GoFish {
	return &GoFish{}
}

func (g *GoFish) Visit() error {
	return g.Request.Do()
}
