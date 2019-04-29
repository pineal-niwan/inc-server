// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

import (
	"github.com/pineal-niwan/busybox/binary"
)

//处理递增
type IncHandler struct {
	*binary.BinaryHandler
}

//反序列化handler,读取字节流到对象中
func NewReadIncHandlerWithOption(data []byte, option *binary.Option) (*IncHandler, error) {
	binHandler, err := binary.NewReadBinaryHandler(data, option)
	if err != nil {
		return nil, err
	} else {
		return &IncHandler{
			BinaryHandler: binHandler,
		}, nil
	}
}

//序列化handler,将对象转化成字节流
func NewWriteIncHandlerWithOption(data []byte, option *binary.Option) (*IncHandler, error) {
	binHandler, err := binary.NewWriteBinaryHandler(data, option)
	if err != nil {
		return nil, err
	} else {
		return &IncHandler{
			BinaryHandler: binHandler,
		}, nil
	}
}

//读取ReqKey
func (p *IncHandler) ReadReqKey() (ret ReqKey, err error) {
	ret.Key, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入ReqKey
func (p *IncHandler) WriteReqKey(v ReqKey) (err error) {
	err = p.WriteString(v.Key)
	if err != nil {
		return
	}
	return
}

//读取ReqKey数组
func (p *IncHandler) ReadReqKeyArray() (ret []ReqKey, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqKey, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqKey()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqKey数组
func (p *IncHandler) WriteReqKeyArray(v []ReqKey) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqKey(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspId
func (p *IncHandler) ReadRspId() (ret RspId, err error) {
	ret.Id, err = p.ReadUint32()
	if err != nil {
		return
	}
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入RspId
func (p *IncHandler) WriteRspId(v RspId) (err error) {
	err = p.WriteUint32(v.Id)
	if err != nil {
		return
	}
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取RspId数组
func (p *IncHandler) ReadRspIdArray() (ret []RspId, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspId, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspId()
		if err != nil {
			return
		}
	}
	return
}

//写入RspId数组
func (p *IncHandler) WriteRspIdArray(v []RspId) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspId(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取ReqKeyList
func (p *IncHandler) ReadReqKeyList() (ret ReqKeyList, err error) {
	ret.KeyList, err = p.ReadStringArray()
	if err != nil {
		return
	}
	return
}

//写入ReqKeyList
func (p *IncHandler) WriteReqKeyList(v ReqKeyList) (err error) {
	err = p.WriteStringArray(v.KeyList)
	if err != nil {
		return
	}
	return
}

//读取ReqKeyList数组
func (p *IncHandler) ReadReqKeyListArray() (ret []ReqKeyList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqKeyList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqKeyList()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqKeyList数组
func (p *IncHandler) WriteReqKeyListArray(v []ReqKeyList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqKeyList(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取KeyIdPair
func (p *IncHandler) ReadKeyIdPair() (ret KeyIdPair, err error) {
	ret.Key, err = p.ReadString()
	if err != nil {
		return
	}
	ret.Id, err = p.ReadUint32()
	if err != nil {
		return
	}
	return
}

//写入KeyIdPair
func (p *IncHandler) WriteKeyIdPair(v KeyIdPair) (err error) {
	err = p.WriteString(v.Key)
	if err != nil {
		return
	}
	err = p.WriteUint32(v.Id)
	if err != nil {
		return
	}
	return
}

//读取KeyIdPair数组
func (p *IncHandler) ReadKeyIdPairArray() (ret []KeyIdPair, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]KeyIdPair, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadKeyIdPair()
		if err != nil {
			return
		}
	}
	return
}

//写入KeyIdPair数组
func (p *IncHandler) WriteKeyIdPairArray(v []KeyIdPair) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteKeyIdPair(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspKeyIdPairList
func (p *IncHandler) ReadRspKeyIdPairList() (ret RspKeyIdPairList, err error) {
	ret.KeyIdPairList, err = p.ReadKeyIdPairArray()
	if err != nil {
		return
	}
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入RspKeyIdPairList
func (p *IncHandler) WriteRspKeyIdPairList(v RspKeyIdPairList) (err error) {
	err = p.WriteKeyIdPairArray(v.KeyIdPairList)
	if err != nil {
		return
	}
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取RspKeyIdPairList数组
func (p *IncHandler) ReadRspKeyIdPairListArray() (ret []RspKeyIdPairList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspKeyIdPairList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspKeyIdPairList()
		if err != nil {
			return
		}
	}
	return
}

//写入RspKeyIdPairList数组
func (p *IncHandler) WriteRspKeyIdPairListArray(v []RspKeyIdPairList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspKeyIdPairList(v[i])
		if err != nil {
			return
		}
	}
	return
}
