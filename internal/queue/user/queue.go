package user

import (
	"fmt"
	"gf-web/internal/queue"
)

type UserLoginQueue struct {

}

func (uq *UserLoginQueue) Execute(payload *queue.QueuePayload) *queue.QueueResult {
	fmt.Println(payload.Body)
	return nil
}
