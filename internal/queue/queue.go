package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

var QM *Queue

type ConsumerInterface interface {
	Execute(payload *QueuePayload) *QueueResult
}

// 队列接口
type IQueue interface {
	Push(payload *QueuePayload) error
	Pop(rcData *RecoverData) error
}

// 消息载体
type QueuePayload struct {
	Topic string `json:"topic"`
	Group string `json:"group"`
	Body interface{} `json:"body"`
}

// 执行结果
type QueueResult struct {
	State bool `json:"state"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// 队列基础服务
type Repo struct {
	Client *gredis.RedisConn
}

// 队列管理器
type Queue struct {
	Repo *Repo
	Handlers map[string]RecoverData
	RecoverCh chan RecoverData
}

type RecoverData struct {
	Topic   string
	Group   string
	Handler ConsumerInterface
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

	h, ok := handler.(ConsumerInterface)
	if !ok {
		panic("handler not realization ConsumerInterface")
	}

	if _, ok := q.Handlers[name]; ok {
		return errors.New("is exits")
	} else {
		q.Handlers[name] = RecoverData{
			Topic: topic,
			Group: group,
			Handler: h,
		}
	}
	return nil
}

func (q *Queue) Push(ctx context.Context, payload *QueuePayload) error {

	payloadStr, _ := json.Marshal(payload)
	name := q.GetQueueName(payload.Topic, payload.Group)
	_, err := q.Repo.Client.Do(ctx, "LPUSH", name, payloadStr)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queue) Pop(ctx context.Context,rcData *RecoverData) error {

	name := q.GetQueueName(rcData.Topic, rcData.Group)
	handler, ok := q.Handlers[name];
	if !ok {
		return errors.New("Execute Not Register")
	}
	//length, err := q.Repo.Client.Do(q.Ctx, "llen", name)
	//if err != nil {
	//	return err
	//}

	reply, err := q.Repo.Client.Do(ctx, "rpop", name)
	if err != nil {
		return err
	}
	resStr := reply.String()
	if len(resStr) > 0 {
		var payload QueuePayload
		err := json.Unmarshal([]byte(resStr), &payload)
		if err != nil {
			return err
		}
		res := handler.Handler.Execute(&payload)
		fmt.Println(res)
	}
	return errors.New("data err")
}

func (q *Queue) SetRecoverCh(ch chan RecoverData) {
	q.RecoverCh = ch
}

func (q *Queue) HandleRecover(ctx context.Context) {
	for  {
		for queueName, handler := range q.Handlers {
			length, err := q.Repo.Client.Do(ctx, "llen", queueName)
			if err != nil {
				fmt.Println(err.Error())
			}
			if length.Int() > 0 {
				q.RecoverCh <- handler
			}
		}
	}
}


func NewQueue(ctx context.Context) *Queue {

	rd, err := g.Redis().Conn(ctx)

	if err != nil {
		panic(err)
	}
	redisRepo := &Repo{
		Client: rd,
	}
	queue := Queue{
		Repo: redisRepo,
		Handlers: make(map[string]RecoverData),
	}

	QM = &queue
	return &queue
}

