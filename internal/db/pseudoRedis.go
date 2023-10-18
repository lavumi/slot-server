package db

import (
	"fmt"
	"time"
)

type PseudoRedis struct{}

var pRedis = PseudoRedis{}
var storage map[string]struct {
	Value      string
	ExpiryTime time.Time
}

type PseudoGetResult struct {
	key string
}

func init() {
	storage = make(map[string]struct {
		Value      string
		ExpiryTime time.Time
	})
}

func (pr *PseudoGetResult) Result() (string, error) {

	data, exists := storage[pr.key]
	if !exists {
		return "", fmt.Errorf("key %s not found", pr.key)
	}
	if time.Now().After(data.ExpiryTime) {
		delete(storage, pr.key)
		return "", fmt.Errorf("data expired")
	}
	return data.Value, nil

}

func (p *PseudoRedis) Get(key string) *PseudoGetResult {
	return &PseudoGetResult{key: key}
}

type PseudoSetResult struct {
}

func (pr *PseudoSetResult) Err() error {
	return nil
}

func (p *PseudoRedis) Set(key string, value string, expiration time.Duration) *PseudoSetResult {
	storage[key] = struct {
		Value      string
		ExpiryTime time.Time
	}{
		Value:      value,
		ExpiryTime: time.Now().Add(expiration),
	}
	return &PseudoSetResult{}
}

type PseudoDelResult struct {
}

func (pr *PseudoDelResult) Result() (int64, error) {
	return 1, nil
}

func (p *PseudoRedis) Del(key string) *PseudoDelResult {
	delete(storage, key)
	return &PseudoDelResult{}
}

func GetRedis() *PseudoRedis {
	return &pRedis
}
