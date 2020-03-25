package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	pb "github.com/kazu69/grpc-kubernetes-microservice-example/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("greeter-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewGreeterClient(conn)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc(("/greet/{name}"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UFT-8")

		vars := mux.Vars(r)
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &pb.HelloRequest{Name: vars["name"]}
		if resp, err := client.SayHello(ctx, req); err == nil {
			msg := fmt.Sprintf("Response is %s", resp.Message)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("GET")

	fmt.Println("Application is running on : 8080 .....")
	http.ListenAndServe(port, routes)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}
