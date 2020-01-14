package main

import (
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestAreElementCached(t *testing.T) {

	sut := newCache(5)

	counter := 0

	f := func(city string) (interface{}, error) {
		counter = counter + 1
		return counter, nil
	}

	first, err := sut.Get("abc", f)

	assert.NilError(t, err)

	second, err := sut.Get("abc", f)

	assert.NilError(t, err)

	assert.Equal(t, first, 1)
	assert.Equal(t, first, second)
}

func TestTimeoutElapsed(t *testing.T) {

	sut := newCache(1)

	counter := 0

	f := func(city string) (interface{}, error) {
		counter = counter + 1
		return counter, nil
	}

	first, err := sut.Get("abc", f)

	assert.NilError(t, err)

	time.Sleep(2 * time.Second)

	second, err := sut.Get("abc", f)

	assert.NilError(t, err)

	assert.Equal(t, first, 1)
	assert.Equal(t, second, 2)
}

func TestRecoverFromPanic(t *testing.T) {

	sut := newCache(1)

	f := func(city string) (interface{}, error) {
		panic("test panic")
	}

	a, err := sut.Get("abc", f)

	assert.Equal(t, a, nil)

	assert.Check(t, err != nil)
	assert.Check(t, strings.Contains(err.Error(), "panic"))

	b, err := sut.Get("abc", f)

	assert.Equal(t, b, nil)

	assert.Check(t, err != nil)

}
