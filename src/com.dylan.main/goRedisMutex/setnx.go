
package main

import (
"fmt"
"sync"
"github.com/go-redis/redis"
	"time"
	"math/rand"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "*****", // no password set
		DB:       0,  // use default DB
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"
	//fmt.Println()
	value:=GetDateTimeFormate()+GetRand17()
	resp := client.SetNX(lockKey, value, time.Second*15)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}
	//
	//// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	if client.Get(lockKey).Val()==value{
		delResp := client.Del(lockKey)
		unlockSuccess, err := delResp.Result()
		if err == nil && unlockSuccess > 0 {
			println("unlock success!")
		} else {
			println("unlock failed", err)
		}
	}else {
		fmt.Println("不是我的锁，不能解锁")
	}
	}

const DATETIMEFORMAT   = "20060102150405"
func GetDateTimeFormate() string {
	return time.Now().In(time.FixedZone("CST", 8*3600)).Format(DATETIMEFORMAT)
}

func GetRand17() string {
	id := fmt.Sprintf("%017x", rand.Int63())
	if len(id) > 17 {
		return id[:17]
	}
	return id
}



func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
