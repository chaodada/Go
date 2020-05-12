package main

import (
	"encoding/json"
	"time"

	"errors"

	"fmt"

	"github.com/go-redis/redis"
)

const (
	URLIDKEY           = "next.url.id"      // 点 分割
	ShortlinkKey       = "shortlink：%s:url" // 点 分割
	URLHashKey         = "urlhash：%s:url"   // 点 分割
	ShortlinkDetailKey = "shortlink：%s:detail"
)

// Redis Cli 包含了一个redis客户端
type RedisCli struct {
	Cli *redis.Client
}

// URL详细信息的结构体
type URLDetail struct {
	URL                 string        `json:"url"`                   // URl
	CreatedAt           string        `json:"created_at"`            // 创建时间
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"` // 过期时间
}

// 初始化redis客户端
func NewRedisCli(add string, pass string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{Addr: add, Password: pass, DB: db}) // 连接数据库
	if _, err := c.Ping().Result(); err != nil {                            // 判断是否连接成功
		panic(err) // 连接失败
	}
	return &RedisCli{Cli: c} // 返回客户端实例
}

// 长地址转化短地址
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {

	// url 转 sha1
	h := toSha1(url)
	// 去缓存中 看看该url 是否存在
	d, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {
		// 说明这个值不存在
	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {
			// 什么都不需要做
		} else {
			return d, nil
		}
	}

	err = r.Cli.Incr(URLIDKEY).Err()
	if err != nil {
		return "", err
	}

	id, err := r.Cli.Get(URLIDKEY).Int64()
	if err != nil {
		return "", err
	}
	eid := base62.EncodeInte64(id)

	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreatedAt:           time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		})
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
}

// 返回详细信息 短地址相关信息
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	d, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, errors.New("没有查询到信息")}
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}
}

// 短地址 转换为长地址
func (r *RedisCli) Unshorten(eid string) (interface{}, error) {
	url, err := r.Cli.Get(fmt.Sprintf(ShortlinkKey, eid)).Result() //拿到长地址
	if err == redis.Nil {
		return "", StatusError{404, errors.New("没有查询到信息")}
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}
}
