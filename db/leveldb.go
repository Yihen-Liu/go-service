package db

import (
	"github.com/Yihen-Liu/go-service/config"
	"github.com/syndtr/goleveldb/leveldb"
)

/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2023-09-27
 */

var LDB *leveldb.DB

func init() {
	var dbErr error
	LDB, dbErr = leveldb.OpenFile(config.CarrierConf.LevelDb, nil)
	if dbErr != nil {
		panic(dbErr)
	}
}
