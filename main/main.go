package main

import (
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"log"
	"os"
	"runtime"
)

func main() {
	//设置GOMAXPROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("init logger error:", err)
		return
	}
	defer logger.Sync()

	appRun := func(c *cli.Context) error {
		return serverRun(c, logger)
	}

	app := cli.App{
		Name:    `IncServer`,
		Usage:   `传入key,获取对应递增的32位整数`,
		Version: `1.0`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  `address`,
				Usage: `rpc service listen address`,
			},
			&cli.StringFlag{
				Name:  `pprofAddress`,
				Usage: `pprof http server address`,
			},
		},
		Action: appRun,
	}

	err = app.Run(os.Args)
	if err != nil {
		logger.Error(app.Name+" exit", zap.Error(err))
		return
	}
}
