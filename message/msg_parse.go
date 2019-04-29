// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

var (
	InitMsgParseHandlerHash map[uint32]fast_rpc.MsgParseHandler
)

func init() {
	var key uint32
	InitMsgParseHandlerHash = make(map[uint32]fast_rpc.MsgParseHandler)
	//生成ReqKey解析函数 command:1 version:0
	key = uint32(1) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqKey{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspId解析函数 command:2 version:0
	key = uint32(2) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspId{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成ReqKeyList解析函数 command:3 version:0
	key = uint32(3) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqKeyList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspKeyIdPairList解析函数 command:4 version:0
	key = uint32(4) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspKeyIdPairList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
}
