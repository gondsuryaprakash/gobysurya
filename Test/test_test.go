package test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type AnonymizeTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *AnonymizeTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatal("Failed to create mock database:", err)
	}
	suite.db = db
	suite.mock = mock
}

func (suite *AnonymizeTestSuite) TearDownTest() {
	suite.db.Close()
}

func (suite *AnonymizeTestSuite) TestAnonymizeTimelineRows() {
	// Create a new context
	ctx := context.Background()

	// Create a test ForgetJob
	job := &ForgetJob{
		TableName:       "test_table",
		EventCategories: []int64{1, 2},
	}

	// Create an instance of the AnonymizeTimelineRows function caller
	dbDao := NewDBDAO(suite.db)

	job.PlayerID = 1
	job.ServiceID = 1
	// expectedQuery := `UPDATE "test_table" FORCE INDEX \(timeline_target_id_service_id_event_category_id_index\) SET target_id = -1, payload = "" WHERE target_id = :player_id AND service_id = :service_id AND event_category_id IN \(:event_categories_0, :event_categories_1\)`
	// suite.mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
	// 	WillReturnResult(sqlmock.NewResult(0, 1))

	categories := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(job.EventCategories)), ","), "([])")
	// expectedQuery := fmt.Sprintf(anonymizeTimelineForPartitionTable, job.TableName, -1, categories)

	// Construct the SQL query with placeholders replaced by actual values
	// query := fmt.Sprintf(`UPDATE %s FORCE INDEX (timeline_target_id_service_id_event_category_id_index) SET target_id = %d, payload = "" WHERE target_id = %d AND service_id = :service_id AND event_category_id IN (%s)`, tableName, playerID, playerID, categories)

	// Construct the regular expression pattern with placeholders replaced by actual values
	pattern := fmt.Sprintf(`UPDATE %s FORCE INDEX \(timeline_target_id_service_id_event_category_id_index\) SET target_id = %d, payload = "" WHERE target_id = %d AND service_id = :service_id AND event_category_id IN \(%s\)`, "tableName", 1, 1, categories)

	// Use regexp.MustCompile to create a regular expression object
	//   regex := regexp.MustCompile(pattern)

	args := []driver.Value{job.PlayerID, job.ServiceID}

	suite.mock.ExpectExec(pattern).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(0, 1))

	caller, err := AnonymizeTimelineRows(&ctx, dbDao, *job)

	// Mock the expected SQL query and its arguments

	// Call the function under test
	rowsAffected, err := caller, err

	// Check if the returned rowsAffected and error are as expected
	suite.NoError(err, "Unexpected error")
	suite.Equal(int64(1), rowsAffected, "Unexpected number of rows affected")

	// Make sure all expectations were met
	suite.NoError(suite.mock.ExpectationsWereMet(), "Not all expectations were met")
}

func TestAnonymizeTestSuite(t *testing.T) {
	suite.Run(t, new(AnonymizeTestSuite))
}
