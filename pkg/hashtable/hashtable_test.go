package hashtable

import (
	"testing"

	"github.com/neelkshah/pandora/config"
)

func TestCreateHashTable(t *testing.T) {
	var response = *CreateHashTable()
	if response.buckets != uint32(config.HASHTABLE_INIT_SIZE) || response.inserts != 0 || response.deletes != 0 {
		t.Fatalf("Fail at values")
	}

	for i := 0; i < int(response.buckets); i++ {
		if response.table[i] == nil {
			t.Fatalf("Fail at buckets")
		}
	}
}
