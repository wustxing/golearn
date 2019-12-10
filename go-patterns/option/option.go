package db

import (
	"fmt"
	"time"
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func Connect(addr string, opts ...Option) error {
	options := options{
		timeout: time.Second,
		caching: true,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	fmt.Println(addr, options)
	return nil
}
