package utility

import (
	"context"
	"fmt"
	"gf-web/internal/model/entity"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"strings"
	"sync"
	"time"
)

type ClickHouse struct {
	client clickhouse.Conn
}

var client *ClickHouse
var once = &sync.Once{}

func NewClickHouse(ctx context.Context) *ClickHouse {
	if client == nil {
		client = &ClickHouse{
			client: nil,
		}
		once.Do(func () {
			var err error
			ck, err := g.Cfg().Get(ctx, "ck")
			if err != nil {
				gerror.New(err.Error())
			}
			config := ck.MapStrVar()
			client.client, err =  clickhouse.Open(&clickhouse.Options{
				Addr: config["address"].Strings(),
				Auth: clickhouse.Auth{
					Database: config["database"].String(),
					Username: config["username"].String(),
					Password: config["password"].String(),
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

	client.client.Ping(ctx)
	return client
}

func (ck *ClickHouse) CreateTable(ctx context.Context, ddl string) error {
	return ck.client.Exec(ctx, ddl)
}

func (ck *ClickHouse) DropTable(ctx context.Context, tableName string) error {
	return ck.client.Exec(ctx, "DROP TABLE IF EXISTS " + tableName)
}

func (ck *ClickHouse) Insert(ctx context.Context, table string, data []interface{}, column []string) error {
	var sql = "INSERT INTO %s (%s) VALUES "
	columnStr := strings.Join(column, ",")
	sql = fmt.Sprintf(sql, table, columnStr)
	batch, err := ck.client.PrepareBatch(ctx, sql)
	if err != nil {
		return err
	}
	for _, v := range data {
		batch.AppendStruct(v)
	}

	return batch.Send()
}

func (ck *ClickHouse) ShowTables(ctx context.Context) (result entity.ShowTables, err error) {
	row := ck.client.QueryRow(ctx, "SHOW TABLES")
	err = row.ScanStruct(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}