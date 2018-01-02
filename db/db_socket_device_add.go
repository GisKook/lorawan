package db

import (
	"github.com/giskook/lorawan/base"
	"context"
)

const (
	SQL_DEVICE_ADD string = "insert into lorawan_device(id) values ($1)"
)

func (db_socket *DBSocket) DeviceAdd(ctx context.Context, id string ,c chan *base.DBResult) {
	_, err := db_socket.db.ExecContext(ctx, SQL_DEVICE_ADD, id)
	if err != nil {
		c <- &base.DBResult{ 
			Err:base.NewErr(err, base.ERR_DEVICE_ADD_CODE, base.ERR_DEVICE_ADD_DESC),
		}

		return
	}

	c<-&base.DBResult{
		Err:base.ERROR_NONE,
	}
}