package leveldb

import (
	"testing"

	"github.com/9z25/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
