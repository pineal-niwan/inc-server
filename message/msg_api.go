// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.
package message

import (
	"context"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

//传入key，获取递增数字
func GetIncNumberByKey(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqKey,
	retryTimes int) (*RspId, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqKey{
			ReqKey: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspId)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspId, nil
}

//传入key和递增的步长，获取递增数字
func GetIncNumberByKeyWithIncStep(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqKeyWithIncNum,
	retryTimes int) (*RspIdWithIncNum, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqKeyWithIncNum{
			ReqKeyWithIncNum: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspIdWithIncNum)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspIdWithIncNum, nil
}

//传入一组key，获取一组对应的递增数字
func GetIncNumberListByKeyList(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqKeyList,
	retryTimes int) (*RspKeyIdPairList, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqKeyList{
			ReqKeyList: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspKeyIdPairList)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspKeyIdPairList, nil
}
