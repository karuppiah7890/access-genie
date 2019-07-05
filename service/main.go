package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	listenAddr := ":8080"

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		fmt.Fprintf(ctx, "{ \"message\": \"pong\" }")
		logger := ctx.Logger()

		logger.Printf("pong!")
	}

	s := fasthttp.Server{
		Name:    "access-genie-service",
		Handler: requestHandler,
	}

	fmt.Println("Listening at port 8080")

	if err := s.ListenAndServe(listenAddr); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
