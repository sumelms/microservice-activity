package tests

import (
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-activity/tests/database"
)

var (
	Now          = time.Now()
	ActivityUUID = uuid.MustParse("d6bb059b-fc2d-485a-9f93-35fc96b1dedd")
	ContentUUID  = uuid.MustParse("05ed8a06-d8e0-4b09-ae83-f9cd80564fa3")
	EmptyRows    = sqlmock.NewRows([]string{})
)

func NewTestDB(queries map[string]string) (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(query)))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)

	return db, mock, sqlStatements
}
