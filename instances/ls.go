package instances

import (
	"github.com/aws/aws-sdk-go/service/ec2"
//	"fmt"
)

func Ls(svc *ec2.EC2){
	resp, err := svc.DescribeInstances(nil)

	if err != nil{
		panic(err)
	}

	collection := new(EC2InstanceCollection)

	for idx, _ := range resp.Reservations {
		for _, inst :=  range resp.Reservations[idx].Instances {
		instance := new(EC2Instance)
		instance.LaunchTime=inst.LaunchTime
		instance.InstanceId=inst.InstanceId
		instance.InstanceType=inst.InstanceType
		instance.KeyName=inst.KeyName
		instance.State=inst.State.Name
		instance.PublicDns=inst.PublicDnsName
		collection.Add(instance)
		}
	}

	collection.Print()
}
