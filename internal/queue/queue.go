package queue

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gredis"
	"strings"
)

type ProductInterface interface {
	Push(topic, group string, message string) error
}

type ConsumerInterface interface {
	Pop(topic, group string) (string, error)
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
	name := q.GetQueueName(topic, group)
	if _, ok := q.Handlers[name]; ok {
		return errors.New("is exits")
	} else {
		q.Handlers[name] = handler
	}
	return nil
}

func (q *Queue) Push(topic, group  string, message string) error {
	name := q.GetQueueName(topic, group)
	_, err := q.Repo.Client.Do(q.Ctx, "LPUSH", name, message)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queue) Pop(topic, group  string) (string, error) {
	name := q.GetQueueName(topic, group)
	reply, err := q.Repo.Client.Do(q.Ctx, "BRPOP", name)
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

