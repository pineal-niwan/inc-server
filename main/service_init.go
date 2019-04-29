package main

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
	"github.com/pineal-niwan/inc-server/handler"
	"github.com/pineal-niwan/inc-server/message"
	"go.uber.org/zap"
	"net"
	"time"
)

func initServiceHandler(ln net.Listener, logger *zap.Logger) (*fast_rpc.Service, error) {
	service := &fast_rpc.Service{}
	option := &fast_rpc.Option{
		Option: &binary.Option{
			DataMaxLen:      1024 * 1024 * 4, //4M
			StringMaxLen:    1024,
			ArrayMaxLen:     1024,
			ExtendExtraSize: 64,
		},

		AcceptDelay:    time.Millisecond,       //1ms
		AcceptMaxDelay: time.Millisecond * 200, //200ms
		AcceptMaxRetry: 1000,

		BufferSize:        1024,
		MaxMsgSize:        1024 * 1024 * 4, //4M
		BufferRecycleSize: 512 * 1024,      //512k
	}

	err := option.Validate()
	if err != nil {
		return nil, err
	}
	service.Init(ln, logger, option, message.InitMsgParseHandlerHash)

	//add your service handler here
	service.AddMsgHandler(&message.MsgReqKey{}, handler.HandleKeyInc)
	service.AddMsgHandler(&message.MsgReqKeyList{}, handler.HandleKeyIncList)

	return service, err
}
