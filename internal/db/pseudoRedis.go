package db

import (
	"fmt"
	"time"
)

type PseudoRedis struct{}

var pRedis PseudoRedis
var storage map[string]string

type PseudoGetResult struct {
	key string
}

func (pr *PseudoGetResult) Result() (string, error) {

	value, exist := storage[pr.key]
	if exist == false {
		return "", fmt.Errorf("key %s not found", pr.key)
	}
	return value, nil
}

func (p *PseudoRedis) Get(key string) *PseudoGetResult {
	return &PseudoGetResult{key: key}
}

type PseudoSetResult struct {
}

func (pr *PseudoSetResult) Err() error {
	return nil
}

func (p *PseudoRedis) Set(key string, value string, time time.Duration) *PseudoSetResult {
	storage[key] = value
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
