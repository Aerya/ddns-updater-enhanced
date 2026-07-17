package data

import (
	"testing"

	"github.com/qdm12/ddns-updater/internal/records"
	"github.com/stretchr/testify/assert"
)

func TestReplaceAll(t *testing.T) {
	t.Parallel()

	db := NewDatabase([]records.Record{{Message: "old"}}, nil)
	want := []records.Record{{Message: "first"}, {Message: "second"}}

	db.ReplaceAll(want)

	assert.Equal(t, want, db.SelectAll())
}
