package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func orderPizza() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxPizzaMakingTimeout, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error")
		}
		defer r.Body.Close()

		fmt.Printf("ordered pizza: %v\n", string(body))
		fmt.Printf("pizza making...\n")

		cooking := make(chan bool)

		c := rand.Intn(100)
		fmt.Println(c)
		if c < 50 {
			// fmt.Println("inside")
			go func() { cooking <- true }()
			// fmt.Println(cooking)
		}

		select {
		case <-cooking:
			str := "Pizza ready, enjoy :)\n"
			fmt.Print(str)
			w.Write([]byte(str))
		case <-r.Context().Done():
			str := "Order has been canceled\n"
			fmt.Print(str)
			w.Write([]byte(str))
		case <-ctxPizzaMakingTimeout.Done():
			str := "Cooking took too much time, client lost\n"
			fmt.Print(str)
			w.Write([]byte(str))
		}
	}
}

func myHandlerRoutes() *http.ServeMux {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", orderPizza())
	return myMux
}

func main() {

	newServer := http.Server{
		Addr:    "localhost:8080",
		Handler: myHandlerRoutes(),
	}

	newServer.ListenAndServe()
	// http.HandleFunc("/", myFuncHandler)
	// http.ListenAndServe("localhost:8080", nil)
}
