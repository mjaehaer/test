package main

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type mainStruct struct {
	redisdb redis.Client
	postgre sqlx.DB
}

// var priority atomic.Uint64

func main() {
	var priority atomic.Uint64
	println("S")
	go func() {
		var mainStruct mainStruct

		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		mainStruct.redisdb = *rdb

		connStr := "user=test password=1234 dbname=test sslmode=disable"
		db, err := sqlx.Open("pgx", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		mainStruct.postgre = *db
	}()
	println(priority.Load())

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	println("S")

	// result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
	// 	"Apple", 72000)
	// if err != nil {
	// 	panic(err)
	// }

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
	// fmt.Printf("%v", mainStruct.redisdb)
}
