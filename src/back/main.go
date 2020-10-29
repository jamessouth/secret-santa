package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jamessouth/secret-santa/src/back/app"
)

func main() {
	reg, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		reg = "us-east-2"
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg),
	})
	logger := aws.NewDefaultLogger()
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		logger.Log("Request: %s/%s, Payload: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})
	if err != nil {
		fmt.Println("bad")
	}
	app.Init(sess)
}
