package utility

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"net"
	"sync"
	"time"
)

type ClickHouse struct {
	Client clickhouse.Conn
}

var client *ClickHouse
var once = &sync.Once{}

func NewClickHouse() *ClickHouse {
	if client == nil {
		client = &ClickHouse{
			Client: nil,
		}
		var err error
		once.Do(func () {
			client.Client, err =  clickhouse.Open(&clickhouse.Options{
				Addr: []string{"127.0.0.1:9005"},
				Auth: clickhouse.Auth{
					Database: "local-test",
					Username: "default",
					Password: "",
				},
				DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
					var d net.Dialer
					return d.DialContext(ctx, "tcp", addr)
				},
				Debug: true,
				Debugf: func(format string, v ...interface{}) {
					fmt.Printf(format + "\n", v)
				},
				Settings: clickhouse.Settings{
					"max_execution_time": 60,
				},
				Compression: &clickhouse.Compression{
					Method: clickhouse.CompressionLZ4,
				},
				DialTimeout:      time.Second * 30,
				MaxOpenConns:     5,
				MaxIdleConns:     5,
				ConnMaxLifetime:  time.Duration(10) * time.Minute,
				ConnOpenStrategy: clickhouse.ConnOpenInOrder,

			})
			if err != nil {
				gerror.New(err.Error())
			}
		})
	}

	client.Client.Ping(context.Background())
	return client
}
