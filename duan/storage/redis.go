package storage

import (
	"duan/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/mattheath/base62"
)

const (
	// URLIDKEY is global counter
	URLIDKEY = "next.url.id"
	// ShortlinkKey mapping the shortlink to the url
	ShortlinkKey = "shortlink:%s:url"
	// URLHashKey mapping the hash of the url to the shortlink
	URLHashKey = "urlhash:%s:url"
	// ShortlinkDetailKey mapping the shortlink to the detail of url
	ShortlinkDetailKey = "shortlink:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

type URLDetail struct {
	URL                 string        `json:"url"`
	CreatedAt           string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

func NewRedisCli(addr, passwd string, db int) *RedisCli {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisCli{Cli: client}
}

// Shorten convert url to shortlink
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {

	fmt.Println("收到的网址", url)

	// 将网址转换为hash    convert url to sha1 hash
	h := utils.ToSha1(url)
	fmt.Println("将网址转换为hash", h)

	// 根据hash 获取缓存
	key := fmt.Sprintf(URLHashKey, h)
	d, err := r.Cli.Get(key).Result()
	if err == redis.Nil {
		// 缓存中不存在
	} else if err != nil {
		return "", err
	} else {
		// 缓存存在 直接返回
		return d, nil
	}

	// increase the global counter
	err = r.Cli.Incr(URLIDKEY).Err() //next.url.id
	if err != nil {
		return "", nil
	}

	// encode global counter to base62
	id, err := r.Cli.Get(URLIDKEY).Int64() //next.url.id

	fmt.Println("id", id)
	if err != nil {
		return "", err
	}
	eid := base62.EncodeInt64(id)
	fmt.Println("eid", eid)

	// store the url against this encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	// store the url against the hash of it
	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	details, err := json.Marshal(&URLDetail{
		URL:                 url,
		CreatedAt:           time.Now().String(),
		ExpirationInMinutes: time.Duration(exp),
	})
	if err != nil {
		return "", err
	}

	// store the url detail against this encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), details, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", nil
	}

	return eid, nil
}

// Shortlink returns the detail of the shortlink
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	d, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", errors.New("unknown short URL")
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}
}

// Unshorten convert shortlink to url
func (r *RedisCli) Unshorten(eid string) (string, error) {
	url, err := r.Cli.Get(fmt.Sprintf(ShortlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", errors.New("unknown short URL")
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}
}
