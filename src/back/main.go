package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jamessouth/secret-santa/src/back/app"
)

func main() {

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("bad")
	}
	logger := aws.NewDefaultLogger()
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		logger.Log("Request: ", r.ClientInfo.ServiceName, r.Operation, "Payload: ", r.Params)
	})
	app.Init(sess)
}
