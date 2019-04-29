package message

import (
	"context"
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
	"testing"
	"time"
)

func TestGetIncNumberByKey(t *testing.T) {
	option := &fast_rpc.CliOption{
		Option: &binary.Option{
			DataMaxLen:      1024 * 1024 * 4, //4M
			StringMaxLen:    1024,            //1024
			ArrayMaxLen:     1024,            //104
			ExtendExtraSize: 64,              //64
		},

		BufferSize:        1024,            //1k
		MaxMsgSize:        1024 * 1024 * 4, //4M
		BufferRecycleSize: 512 * 1024,      //512k
		RetreatTime:       200 * time.Millisecond, //200ms
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cli, err := fast_rpc.NewCli(ctx, "localhost:20001", 5, option, InitMsgParseHandlerHash)
	if err != nil {
		t.Errorf("connect error:%+v", err)
		return
	}

	rspId, err := GetIncNumberByKey(ctx, cli, ReqKey{Key: "1"}, 3)
	if err != nil {
		t.Errorf("get key error:%+v", err)
		t.Errorf("client max msgsize:%+v", cli.MaxMsgSize)
		return
	}

	t.Logf("rspId:%+v", rspId)
}
