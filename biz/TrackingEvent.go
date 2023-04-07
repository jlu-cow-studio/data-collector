package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
)

func RecordEvent(event *mysql_model.Event) error {
	return mysql.GetDBConn().Table("tracking_events").Create(event).Error
}
