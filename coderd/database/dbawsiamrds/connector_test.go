package dbawsiamrds_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/coder/coder/v2/coderd/database/dbawsiamrds"
	"github.com/coder/coder/v2/testutil"
)

func TestConnector(t *testing.T) {
	url := os.Getenv("DBAWSIAMRDS_TEST_URL")
	if url == "" {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), testutil.WaitShort)
	defer cancel()

	db, err := dbawsiamrds.NewDB(ctx, url)
	require.NoError(t, err)

	var i int
	err = db.SelectContext(ctx, &i, "select 1;")
	require.NoError(t, err)
	assert.Equal(t, 1, i)
}
