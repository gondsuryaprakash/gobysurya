package cache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func GoCache() {

	// Creating new cache with 5 min expiry time and the 10 min purgs time
	c := cache.New(time.Minute*5, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

}
