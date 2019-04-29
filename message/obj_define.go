// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

//获取递增
type ReqKey struct {
	//获取递增
	Key string
}

//返回递增的值
type RspId struct {
	//返回递增
	Id uint32
	//获取过程是否有err,如果有,则此字符串表示error内容
	Err string
}

//获取递增列表
type ReqKeyList struct {
	//获取递增列表
	KeyList []string
}

//key - uint32对
type KeyIdPair struct {
	//key
	Key string
	//递增id
	Id uint32
}

//返回递增的值
type RspKeyIdPairList struct {
	//返回递增列表 (key-id)的列表
	KeyIdPairList []KeyIdPair
	//获取过程是否有err,如果有,则此字符串表示error内容
	Err string
}
