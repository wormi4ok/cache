package cache

import "errors"

var ErrorNoValue = errors.New("key does not exist")

type Cache interface {
	Write(key string, value interface{})
	Read(key string) (interface{}, error)
}
