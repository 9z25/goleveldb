package iterator_test

import (
	"testing"

	"github.com/9z25/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
