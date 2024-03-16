package db

import (
	"fmt"
	"github.com/calmw/fdb"
	"home_assistant/log"
)

func InitFdb() {
	log.Logger.Info("Init fdb db")
	// 打开fdb redis
	var err error
	opts := fdb.DefaultOption
	opts.DataFileSize = 60 * 1024 * 1024
	FDB, err = fdb.Open(opts)
	if err != nil {
		if err != nil {
			panic(fmt.Sprintf("Init fdb db failed:%v", err))
		}
	}
	log.Logger.Info("Init fdb success")
}
