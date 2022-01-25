package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteRead(t *testing.T) {
	want := "value1"
	cache := NewLRU(1)

	cache.Write("test1", want)
	got, err := cache.Read("test1")

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestReadNonExistent(t *testing.T) {
	cache := NewLRU(1)

	got, err := cache.Read("test1")

	assert.ErrorIs(t, err, ErrorNoValue)
	assert.Equal(t, "", got)
}

func TestWriteOverflow(t *testing.T) {

	cache := NewLRU(1)

	cache.Write("test1", "value1")
	cache.Write("test2", "value2")

	_, err := cache.Read("test1")
	assert.ErrorIs(t, err, ErrorNoValue)

	got, err := cache.Read("test2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", got)
}

func TestLRUEviction(t *testing.T) {
	cache := NewLRU(2)

	cache.Write("test1", "value1")
	cache.Write("test2", "value2")

	_, _ = cache.Read("test1")

	cache.Write("overflowKey", "overflowValue")
	_, err := cache.Read("test2")
	assert.ErrorIs(t, err, ErrorNoValue)

	_, err = cache.Read("test1")
	assert.NoError(t, err)
}

func BenchmarkCacheWrite(b *testing.B) {
	b.ReportAllocs()
	cache := NewLRU(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Write("test", 123)
	}
}

var valTmp interface{}

func BenchmarkCacheRead(b *testing.B) {
	b.ReportAllocs()
	cache := NewLRU(1)
	cache.Write("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		valTmp, _ = cache.Read("key")
	}
	assert.Equal(b, valTmp.(string), "value")
}
