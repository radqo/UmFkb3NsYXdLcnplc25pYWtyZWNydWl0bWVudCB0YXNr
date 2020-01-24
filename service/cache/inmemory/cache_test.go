package inmemory

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestAreElementCached(t *testing.T) {

	sut := New(5)

	_, found := sut.Get("abc")

	assert.Equal(t, found, false)

	sut.Set("abc", 123)
	m, found := sut.Get("abc")

	assert.Equal(t, found, true)
	assert.Equal(t, m, 123)
}

func TestTimeoutElapsed(t *testing.T) {

	sut := New(1)

	sut.Set("abc", 123)

	_, found := sut.Get("abc")

	assert.Equal(t, found, true)

	time.Sleep(2 * time.Second)

	_, found = sut.Get("abc")

	assert.Equal(t, found, false)

}
