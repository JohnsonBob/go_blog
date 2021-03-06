package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_blog/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

func SetUp() {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.Config.Redis.MaxIdle,
		MaxActive:   setting.Config.Redis.MaxActive,
		IdleTimeout: setting.Config.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", setting.Config.Redis.Host)
			if err != nil {
				return nil, err
			}
			if setting.Config.Redis.Password != "" {
				_, err := dial.Do("AUTH", setting.Config.Redis.Password)
				if err != nil {
					_ = dial.Close()
					return nil, err
				}
			}
			return dial, err
		},
		TestOnBorrow: func(conn redis.Conn, time time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) { _ = conn.Close() }(conn)
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func IsExists(key string) bool {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) { _ = conn.Close() }(conn)
	isExists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return isExists

}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func() { _ = conn.Close() }()
	reply, err := conn.Do("GET", key)
	if reply == nil && err == nil {
		return nil, nil
	}
	bytes, err := redis.Bytes(reply, err)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func() { _ = conn.Close() }()
	reply, err := conn.Do("DEL", key)
	if reply == nil && err == nil {
		return false, nil
	}
	return redis.Bool(reply, err)
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer func() { _ = conn.Close() }()
	reply, err := conn.Do("KEYS", "*"+key+"*")
	if reply == nil && err == nil {
		return nil
	}
	keys, err := redis.Strings(reply, err)
	if err != nil {
		return err
	}
	for _, value := range keys {
		_, err := Delete(value)
		if err != nil {
			return err
		}
	}
	return nil
}
