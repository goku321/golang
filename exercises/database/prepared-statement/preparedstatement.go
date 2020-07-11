package main

import (
	"database/sql"
	"os"

	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var stmtCache *sq.StmtCache

func runQuery(siteID, triggerID int, customerUserID string) (sql.NullFloat64, error) {
	stmtBuilder := sq.StatementBuilder.RunWith(stmtCache).PlaceholderFormat(sq.Dollar)
	sqlString, args, err := stmtBuilder.
		Select("survey_rating").
		From("survey_responses").
		Where(sq.Eq{
			"site_id":          siteID,
			"customer_user_id": customerUserID,
			"trigger_id":       triggerID,
		}).
		OrderBy("id DESC").
		Limit(1).
		ToSql()
	if err != nil {
		return sql.NullFloat64{}, err
	}
	var rating sql.NullFloat64
	row := stmtCache.QueryRow(sqlString, args...)
	err = row.Scan(&rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return rating, nil
		}
		return rating, err
	}

	return rating, nil
}

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("PG_TEST_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	// Init statement cache.
	stmtCache = sq.NewStmtCache(db)
	res, err := runQuery(606078601, 22267, "5217548")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
