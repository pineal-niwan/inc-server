package main

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
	"github.com/pineal-niwan/inc-server/inc_hash"
	"github.com/pineal-niwan/inc-server/message"
	"go.uber.org/zap"
	"net"
	"time"
)

//递增服务
type IncService struct {
	*fast_rpc.Service
	numIncHash *inc_hash.NumIncHash
}

//初始化服务
func initServiceHandler(ln net.Listener, logger *zap.Logger) (*IncService, error) {
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

	rpcService := &fast_rpc.Service{}
	rpcService.Init(ln, logger, option, message.InitMsgParseHandlerHash)

	service := &IncService{
		Service:    rpcService,
		numIncHash: &inc_hash.NumIncHash{},
	}

	//add your service handler here
	service.AddMsgHandler(&message.MsgReqKey{}, service.HandleKeyInc)
	service.AddMsgHandler(&message.MsgReqKeyWithIncNum{}, service.HandleKeyIncWithStep)
	service.AddMsgHandler(&message.MsgReqKeyList{}, service.HandleKeyIncList)
	service.numIncHash.Init(service.getIncNumberByKey)

	return service, err
}

//由单个key获取其对应的id
func (s *IncService) HandleKeyInc(inMsg fast_rpc.IMsg) (outMsg fast_rpc.IMsg, err error) {
	req, ok := inMsg.(*message.MsgReqKey)
	if !ok {
		err = fast_rpc.ErrNotExpectMsg
		return
	}
	id, fErr := s.handleKeyInc(req.Key, 1)

	rsp := &message.MsgRspId{}
	if fErr != nil {
		rsp.Err = fErr.Error()
	} else {
		rsp.Id = id
	}
	return rsp, nil
}

//由单个key以及指定的步长获取其对应的id
func (s *IncService) HandleKeyIncWithStep(inMsg fast_rpc.IMsg) (outMsg fast_rpc.IMsg, err error) {
	req, ok := inMsg.(*message.MsgReqKeyWithIncNum)
	if !ok {
		err = fast_rpc.ErrNotExpectMsg
		return
	}
	id, fErr := s.handleKeyInc(req.Key, req.IncNum)

	rsp := &message.MsgRspIdWithIncNum{}
	if fErr != nil {
		rsp.Err = fErr.Error()
	} else {
		rsp.Id = id
		rsp.IncNum = req.IncNum
	}
	return rsp, nil
}

//由多个key获取其对应的id
func (s *IncService) HandleKeyIncList(inMsg fast_rpc.IMsg) (outMsg fast_rpc.IMsg, err error) {
	req, ok := inMsg.(*message.MsgReqKeyList)
	if !ok {
		err = fast_rpc.ErrNotExpectMsg
		return
	}

	keyIdPairList := make([]message.KeyIdPair, 0, len(req.KeyList))
	var rspErr string
	for _, key := range req.KeyList {
		id, fErr := s.handleKeyInc(key, 1)
		if fErr != nil {
			rspErr = fErr.Error()
		} else {
			keyIdPairList = append(keyIdPairList, message.KeyIdPair{
				Id:  id,
				Key: key,
			})
		}
	}

	rsp := &message.MsgRspKeyIdPairList{}
	rsp.Err = rspErr
	rsp.KeyIdPairList = keyIdPairList
	return rsp, nil
	return
}

func (s *IncService) handleKeyInc(key string, step uint32) (number uint32, err error) {
	return s.numIncHash.Get(key, step)
}
