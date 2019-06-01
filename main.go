package main

import (
	
	"go-web/controller"
	"net/http"

	//"github.com/go-redis/redis"
	
	"log"
)

// func httpHandle(w http.ResponseWriter, req *http.Request) {
// 	//ret := core.Succeed(w);
// 	fmt.Println(ret)
// 	fmt.Println("path", req.URL.Path)
// 	fmt.Fprintf(w, req.URL.Path)
// }

func main() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// })
	
	//test(w http.ResponseWriter)
	// conf := config.ReadConfig();
    // log.Println("conf", conf)
	// http.HandleFunc("/", httpHandle)
	 http.HandleFunc("/user/login", controller.UserLogin)
	 err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Listen", err)
	}
}


