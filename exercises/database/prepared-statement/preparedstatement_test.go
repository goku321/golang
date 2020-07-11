package main

import (
	"os"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkRunQueryWithCache(b *testing.B) {
	db, err := sqlx.Open("postgres", os.Getenv("PG_TEST_DATABASE"))
	require.NoError(b, err)
	stmtCache = sq.NewStmtCache(db)

	for i := 0; i < b.N; i++ {
		_, err := runQueryWithCache(606078601, 22267, "5217548")
		assert.NoError(b, err)
	}
}

func BenchmarkRunQuery(b *testing.B) {
	db, err := sqlx.Open("postgres", os.Getenv("PG_TEST_DATABASE"))
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := runQuery(db, 606078601, 22267, "5217548")
		assert.NoError(b, err)
	}
}


