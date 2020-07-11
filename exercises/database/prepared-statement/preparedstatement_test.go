package main

import (
	"context"
	"os"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkRunQueryWithCache(b *testing.B) {
	db, err := sqlx.Open("postgres", os.Getenv("PG_TEST_DATABASE"))
	require.NoError(b, err)
	defer db.Close()
	stmtCache = sq.NewStmtCache(db)

	for i := 0; i < b.N; i++ {
		_, err := runQueryWithCache(606078601, 22267, "5217548")
		assert.NoError(b, err)
	}
}

func BenchmarkRunQuery(b *testing.B) {
	db, err := sqlx.Open("postgres", os.Getenv("PG_TEST_DATABASE"))
	require.NoError(b, err)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		_, err := runQuery(db, 606078601, 22267, "5217548")
		assert.NoError(b, err)
	}
}

func BenchmarkRunQueryWithPGX(b *testing.B) {
	db, err := pgx.Connect(context.Background(), os.Getenv("PG_TEST_DATABASE"))
	require.NoError(b, err)
	defer db.Close(context.Background())

	for i := 0; i < b.N; i++ {
		_, err := runQueryWithPGX(db, 606078601, 22267, "5217548")
		assert.NoError(b, err)
	}
}
