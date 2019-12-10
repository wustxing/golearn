package db

import (
	"testing"
	"time"
)

func TestOption(t *testing.T) {
	Connect("127.0.0.1:8080", WithCaching(false), WithTimeout(time.Minute))
	Connect("127.0.0.1:8080")
}
