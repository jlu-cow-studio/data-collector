package handler

import "github.com/jlu-cow-studio/common/dal/rpc/data_collector"

type Handler struct {
	data_collector.UnimplementedDataCollectorServiceServer
}
