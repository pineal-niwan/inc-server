// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.
package message

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

//获取递增
type MsgReqKey struct {
	ReqKey
}

//获取命令行
func (msg *MsgReqKey) GetCmd() uint16 {
	return uint16(1)
}

//获取版本号
func (msg *MsgReqKey) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqKey) GetCode() uint32 {
	return uint32(1) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqKey) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqKey(msg.ReqKey)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(1),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqKey) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqKey, err = reader.ReadReqKey()
	return err
}

//返回递增的值
type MsgRspId struct {
	RspId
}

//获取命令行
func (msg *MsgRspId) GetCmd() uint16 {
	return uint16(2)
}

//获取版本号
func (msg *MsgRspId) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspId) GetCode() uint32 {
	return uint32(2) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspId) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspId(msg.RspId)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(2),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspId) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspId, err = reader.ReadRspId()
	return err
}

//获取递增-带增加的数字
type MsgReqKeyWithIncNum struct {
	ReqKeyWithIncNum
}

//获取命令行
func (msg *MsgReqKeyWithIncNum) GetCmd() uint16 {
	return uint16(3)
}

//获取版本号
func (msg *MsgReqKeyWithIncNum) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqKeyWithIncNum) GetCode() uint32 {
	return uint32(3) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqKeyWithIncNum) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqKeyWithIncNum(msg.ReqKeyWithIncNum)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(3),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqKeyWithIncNum) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqKeyWithIncNum, err = reader.ReadReqKeyWithIncNum()
	return err
}

//返回递增的值
type MsgRspIdWithIncNum struct {
	RspIdWithIncNum
}

//获取命令行
func (msg *MsgRspIdWithIncNum) GetCmd() uint16 {
	return uint16(4)
}

//获取版本号
func (msg *MsgRspIdWithIncNum) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspIdWithIncNum) GetCode() uint32 {
	return uint32(4) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspIdWithIncNum) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspIdWithIncNum(msg.RspIdWithIncNum)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(4),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspIdWithIncNum) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspIdWithIncNum, err = reader.ReadRspIdWithIncNum()
	return err
}

//获取递增列表
type MsgReqKeyList struct {
	ReqKeyList
}

//获取命令行
func (msg *MsgReqKeyList) GetCmd() uint16 {
	return uint16(5)
}

//获取版本号
func (msg *MsgReqKeyList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqKeyList) GetCode() uint32 {
	return uint32(5) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqKeyList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqKeyList(msg.ReqKeyList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(5),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqKeyList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqKeyList, err = reader.ReadReqKeyList()
	return err
}

//返回递增的值
type MsgRspKeyIdPairList struct {
	RspKeyIdPairList
}

//获取命令行
func (msg *MsgRspKeyIdPairList) GetCmd() uint16 {
	return uint16(6)
}

//获取版本号
func (msg *MsgRspKeyIdPairList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspKeyIdPairList) GetCode() uint32 {
	return uint32(6) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspKeyIdPairList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspKeyIdPairList(msg.RspKeyIdPairList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(6),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspKeyIdPairList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspKeyIdPairList, err = reader.ReadRspKeyIdPairList()
	return err
}
