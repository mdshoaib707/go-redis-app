package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"os"

	"fmt"
	"reflect"
	"github.com/gomodule/redigo/redis"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	mymessage := "Hello, you have "+ cookie.Value +" visitors on this page"

	// io.WriteString(res, cookie.Value)
	io.WriteString(res, mymessage)

	fmt.Println(reflect.TypeOf(count))

	fmt.Println("Page hits is:::")

	fmt.Println(count)

	// Read/Write to Redis
	redishost := os.Getenv("SERVER_REDIS")
	redisport := os.Getenv("SERVER_REDIS_PORT")
	finalredis := redishost + ":" + redisport

	fmt.Println("redis details:",finalredis)

	conn, err := redis.Dial("tcp", finalredis)
	if err != nil {
		log.Fatal(err)
	}
	// properly closed before exiting the main() function.
	defer conn.Close()

	_, err = conn.Do("SET", "pagehits", count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added pagecount in REDIS!")

	strs, err := redis.String(conn.Do("GET", "pagehits"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strs)
}
