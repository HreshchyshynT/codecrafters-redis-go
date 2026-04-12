package store

import (
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

type Data struct {
	Value     resp.Value
	CreatedAt time.Time
	ExpireIn  time.Duration
}

func NewData(value resp.Value, expireIn ...time.Duration) Data {
	var expire time.Duration
	if len(expireIn) > 0 {
		expire = expireIn[0]
	}

	return Data{
		Value:     value,
		CreatedAt: time.Now(),
		ExpireIn:  expire,
	}
}

func (d Data) isExpired() bool {
	return d.ExpireIn > 0 && d.CreatedAt.Add(d.ExpireIn).Before(time.Now())
}
