package misc

import (
	"fmt"
	"time"
	
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	
	defer c.Close()
	
	v, err := c.Do("set", "name", "ethan", "EX", "5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	
	v1, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v1)
	
	time.Sleep(time.Second * 8)
	v2, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v2)
		
	}
	
}
