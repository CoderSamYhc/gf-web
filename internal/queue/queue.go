package queue

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"strings"
)

type ProductInterface interface {
	Push(queueName string, message string) error
}

type ConsumerInterface interface {
	Pop(queueName string) (string, error)
}

type IQueue interface {
	ProductInterface
	ConsumerInterface
}



type Repo struct {
	Client *gredis.RedisConn
}

type Queue struct {
	Repo *Repo
	Ctx context.Context
	Handlers map[string]interface{}
}

func (q *Queue) GetQueueName(topic, group string) string {
	var name string
	if len(group) > 0 {
		name = strings.Replace("Queue_%t::%g", "%t", topic, -1)
		name = strings.Replace(name, "%g", group, -1)
	} else {
		name = strings.Replace("Queue_%t", "%t", topic, -1)
	}
	return name
}

func (q *Queue) RegisterQueue(topic, group string, handler interface{}) error {
	return nil
}

func (q *Queue) Push(queueName string, message string) error {
	_, err := q.Repo.Client.Do(q.Ctx, "LPUSH", queueName, message)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queue) Pop(queueName string) (string, error) {
	reply, err := q.Repo.Client.Do(q.Ctx, "LPOP", queueName)
	if err != nil {
		return "", err
	}
	return reply.String(), nil
}

func NewQueue(rd *gredis.RedisConn, ctx context.Context) Queue {

	redisRepo := &Repo{
		Client: rd,
	}
	queue := Queue{
		Repo: redisRepo,
		Ctx: ctx,
	}

	return queue
}

