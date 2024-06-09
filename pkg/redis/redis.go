package redis

import (
	"fmt"
	"sync"

	"github.com/redis/rueidis"
)

var redisInstance rueidis.Client
var redisOnce sync.Once

func NewRedissClient() rueidis.Client {
	if redisInstance == nil {
		redisOnce.Do(func() {

			c, err := rueidis.NewClient(rueidis.ClientOption{
				InitAddress:  []string{"127.0.0.1:6379"},
				DisableCache: true,
			})
			if err != nil {
				fmt.Println(err)
			}
			redisInstance = c

		})
	}
	return redisInstance
}

/*
InitAddress:  []string{cfg.Redis.Queue.Host + ":" + cfg.Redis.Queue.Port},
		Username:     cfg.Redis.Queue.Username,
		Password:     cfg.Redis.Queue.Password,
		SelectDB:     database,
		DisableCache: true,

*/
