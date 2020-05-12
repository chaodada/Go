package main

import (
	"log"
	"os"
	"strconv"

	"fmt"
)

type Env struct {
	S Storage
}

//func getEnv() *Env {
func getEnv() {
	// redis 连接地址
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	// redis 密码
	pass := os.Getenv("APP_REDIS_PASSWD")
	if pass == "" {
		pass = ""
	}
	// redis 数据库
	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = "0"
	}

	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("连接redis(addr: %s password: %s db: %d)", addr, pass, db)

	r := NewRedisCli(addr, pass, db)

	fmt.Println(r)
	//return &Env{S: r}
}
