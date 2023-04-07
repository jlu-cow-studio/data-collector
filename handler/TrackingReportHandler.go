package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/data_collector"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	redis_model "github.com/jlu-cow-studio/common/model/dao_struct/redis"
	"github.com/jlu-cow-studio/data-collector/biz"
)

func (h *Handler) TrackingReport(ctx context.Context, req *data_collector.TrackingReportReq) (res *data_collector.TrackingReportRes, err error) {
	res = &data_collector.TrackingReportRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	cmd := redis.DB.Get(redis.GetUserTokenKey(req.Base.Token))
	if cmd.Err() != nil {
		res.Base.Message = cmd.Err().Error()
		res.Base.Code = "401"
		return res, nil
	}

	info := &redis_model.UserInfo{}

	if err := json.Unmarshal([]byte(cmd.Val()), info); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "402"
		return res, nil
	}

	uid, _ := strconv.Atoi(info.Uid)

	if err = biz.RecordEvent(&mysql_model.Event{
		ItemID:    uint(req.ItemId),
		Timestamp: time.Now(),
		EventType: req.Behavior,
		UserID:    uint(uid),
	}); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	}

	res.Base.Message = ""
	res.Base.Code = "200"

	return
}
