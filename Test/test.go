package test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
)

func IsEven(n int) bool {
	if n%2 != 0 {
		return false
	}

	return true
}

const TimelineTableName = "timeline"
const Timeline2TableName = "timeline_2"

const anonymizeTimeline = `
    UPDATE %s FORCE INDEX (timeline_target_id_service_id_event_category_id_index) SET target_id = %d, payload = "" where target_id = :player_id and service_id = :service_id and event_category_id in (%s)
`

const anonymizeTimelineForPartitionTable = `
    UPDATE %s FORCE INDEX (main_idx) SET player_id = %d, payload = "" WHERE player_id = :player_id AND service_id = :service_id AND event_category_id IN (%s)
`

func AnonymizeTimelineRows(cx *context.Context, dao DaoDB, job ForgetJob) (int64, error) {
	categories := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(job.EventCategories)), ","), "([])")
	query := fmt.Sprintf(anonymizeTimeline, job.TableName, -1, categories)

	if job.TableName != TimelineTableName && job.TableName != Timeline2TableName {
		query = fmt.Sprintf(anonymizeTimelineForPartitionTable, job.TableName, -1, categories)
	}

	// log := ctx.GetSharedLogger(*cx)
	// log.Infof("Anonymize:  %s", query)

	rows, err := dao.PerformAnonymizeTimeline(cx, query, job)
	if err != nil {
		return rows, errors.New("failed to anonymize data")
	}

	return rows, nil
}

type ForgetJob struct {
	ID                string          `json:"id" db:"job_id"`
	ForgetRequestID   string          `json:"forget_request_id" db:"request_id"`
	CallbackURL       string          `json:"callback_url"`
	CallbackToken     string          `json:"callback_token"`
	TableName         string          `json:"table_name"`
	NumRecordsUpdated int64           `json:"num_records_updated"`
	Status            ForgetJobStatus `json:"status" db:"status"`
	ErrorMessage      string          `json:"error_message"`
	ServiceID         int64           `json:"service_id" db:"service_id"`
	PlayerID          int64           `json:"player_id" db:"player_id"`
	EventCategories   []int64         `json:"event_categories" db:"event_categories"`
	StartTime         int64           `json:"start_time"`
	EndTime           int64           `json:"end_time"`
}

type ForgetJobStatus int64

type DaoDB interface {
	PerformAnonymizeTimeline(cx *context.Context, query string, job ForgetJob) (int64, error)
}

func NewDBDAO(db *sql.DB) DaoDB {
	return &DBDao{
		DB: sqlx.NewDb(db, "sqlmock"),
	}
}

// PerformAnonymizeTimeline executes an anonymize query on the Timeline DB
func (dao *DBDao) PerformAnonymizeTimeline(cx *context.Context, query string, job ForgetJob) (int64, error) {
	result, err := dao.DB.NamedExec(query, job)
	if err != nil {
		// log := ctx.GetSharedLogger(*cx)
		// log.Errorf("Anonymize Failed: ((%v)) %v", job, err.Error())

		fmt.Println(err)
		return 0, err
	}

	return result.RowsAffected()
}

type DBDao struct {
	DB *sqlx.DB
	// cache      *cache.Cache
	// Config     dbconfig.Config
	// traceAttrs []attribute.KeyValue
	mu *sync.Mutex
}
