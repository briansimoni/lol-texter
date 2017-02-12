package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"fmt"
)

func text () {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		panic(err)
	}


	// Create a SNS client from just a session.
	svc := sns.New(sess)
	err = sendToAWS(PhoneNumber, "Austin is playing League!!!", svc)
	if err != nil {
		fmt.Println("Failed to send to Brian", err.Error())
	}

	err = sendToAWS(PhoneNumber2, "You have notified Brian you are playing League!", svc)
	if err != nil {
		fmt.Println("Failed to send to Austin", err.Error())
	}
}


func sendToAWS(phoneNumber string, message string, svc *sns.SNS) error {
	params := &sns.PublishInput{
		Message: aws.String(message),
		PhoneNumber:      aws.String(phoneNumber),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		return err
	}

	// Pretty-print the response data.
	fmt.Println("the message should've sent", resp)
	return nil
}