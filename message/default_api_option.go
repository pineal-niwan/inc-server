package message

import (
	"context"
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
	"time"
)

//缺省的客户端参数
func DefualtCliOption() *fast_rpc.CliOption {
	return &fast_rpc.CliOption{
		Option: &binary.Option{
			DataMaxLen:      1024 * 1024 * 4, //4M
			StringMaxLen:    1024,            //1024
			ArrayMaxLen:     1024,            //1024
			ExtendExtraSize: 64,              //64
		},

		BufferSize:        128,                    //128
		MaxMsgSize:        1024 * 1024 * 4,        //4M
		BufferRecycleSize: 512 * 1024,             //512k
		RetreatTime:       200 * time.Millisecond, //200ms
	}
}

//新建缺省客户端
func NewDefaultCli(ctx context.Context, url string) (*fast_rpc.Cli, error) {
	cli, err := fast_rpc.NewCli(ctx, url, 100, DefualtCliOption(), InitMsgParseHandlerHash)
	return cli, err
}
