package main

import (
	"gopkg.in/redsync.v1"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

func main(){
	TestRedis()
}

//redis命令执行函数
func DoRedisCmdByConn(conn *redis.Pool,commandName string, args ...interface{}) (interface{}, error) {
	redisConn := conn.Get()
	defer redisConn.Close()
	//检查与redis的连接
	return redisConn.Do(commandName, args...)
}


func TestRedis() {

	//单个锁
	pool := newRedisPool()
	rs := redsync.New([]redsync.Pool{})
	mutex1 := rs.NewMutex("test-redsync1")

	mutex1.Lock()
	conn1 := pool.Get()
	conn1.Do("SET","name1","ywb1")
	re,err:=conn1.Do("get","name1")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(re.([]byte)))
	conn1.Close()
	conn2 := pool.Get()
	fmt.Println(conn2.Do("SET","name1","ywb1"))
	mutex1.Unlock()

}

func newTestMutexes(pools []redsync.Pool, name string, n int) []*redsync.Mutex {
	mutexes := []*redsync.Mutex{}
	for i := 0; i < n; i++ {
		mutexes = append(mutexes,redsync.New(pools).NewMutex(name,
			redsync.SetExpiry(time.Duration(2)*time.Second),
			redsync.SetRetryDelay(time.Duration(10)*time.Millisecond)),
		)
	}
	return mutexes
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: time.Duration(24) * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "49.234.95.92:6379")
			if err != nil {
				panic(err.Error())
				//s.Log.Errorf("redis", "load redis redisServer err, %s", err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				//s.Log.Errorf("redis", "ping redis redisServer err, %s", err.Error())
				return err
			}
			return err
		},
	}
}

func newRedisPool()*redis.Pool{
	return &redis.Pool{
		MaxIdle:50,
		MaxActive:30,
		IdleTimeout:300*time.Second,
		Dial: func() (redis.Conn,  error) {
			c,err := redis.Dial("tcp","49.234.95.92:6379")
			if err != nil {
				fmt.Println(err)
				return nil,err
			}
			//2、访问认证
			if _, err =c.Do("AUTH","muwenbo");err!=nil{
				c.Close()
				return nil,err
			}
			return c,nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t)<time.Minute{
				return nil
			}
			_,err := conn.Do("PING")
			return err
		},
	}
}

