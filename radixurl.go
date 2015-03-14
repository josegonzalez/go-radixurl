package radixurl

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"net/url"
	"os"
	"strings"
	"time"
)

func Connect() (*redis.Client, error) {
	return ConnectToURL(os.Getenv("REDIS_URL"))
}

func ConnectToURL(s string) (c *redis.Client, err error) {
	redisUrl, err := url.Parse(s)
	if err != nil {
		return
	}

	splitted := strings.Split(redisUrl.Host, ":")
	host, port := splitted[0], "6379"

	switch {
	case len(splitted) > 2:
		err = fmt.Errorf("parse error %q in HOST:PORT", redisUrl.Host)
	case len(splitted) > 1:
		host, port = splitted[0], splitted[1]
	default:
		host, port = splitted[0], "6379"
	}

	if err != nil {
		return
	}

	database := ""

	if len(redisUrl.Path) > 1 {
		database = strings.TrimPrefix(redisUrl.Path, "/")
	}

	c, err = redis.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Duration(10)*time.Second)
	if err == nil && database != "" {
		r := c.Cmd("select", database)
		if r.Err != nil {
			return
		}
	}
	return c, err
}
