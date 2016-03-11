package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)


func main() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

	resp, err := svc.DescribeInstances(nil)
	if err != nil{
		panic(err)
	}

	fmt.Println("> Number of reservation sets: ",len(resp.Reservations))
	for idx, res := range resp.Reservations {
		fmt.Println(" > Number of instances: ", len(res.Instances))
		for _, inst :=  range resp.Reservations[idx].Instances {
			fmt.Println(*inst)
		}
	}
}
