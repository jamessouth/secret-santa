package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func showTime() time.Time {
	return time.Now()
}

func main() {

	// sess, err := session.NewSession()
	// if err != nil {
	// 	fmt.Println("bad")
	// }
	// logger := aws.NewDefaultLogger()
	// sess.Handlers.Send.PushFront(func(r *request.Request) {
	// 	// Log every request made and its payload
	// 	logger.Log("Request: ", r.ClientInfo.ServiceName, r.Operation, "Payload: ", r.Params)
	// })
	// app.Init(sess)

	lambda.Start(showTime)
}
