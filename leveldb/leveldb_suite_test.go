package leveldb

import (
	"testing"

	"github.com/seopub/btcsuite_goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
