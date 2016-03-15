package instances

import(
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

func Run(svc *ec2.EC2, ImageId *string, count int, keyPair *string, instanceType *string) {
	c := int64(count)

	params := &ec2.RunInstancesInput{
		ImageId: ImageId,
		MaxCount: &c,
		MinCount: &c,
		InstanceType: instanceType,
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination: aws.Bool(true),
				DeviceIndex: aws.Int64(0),
			},
		},
		KeyName: keyPair,
	}

	resp,err := svc.RunInstances(params)

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	collection := new(EC2InstanceCollection)
	for _, inst := range resp.Instances {
		instance := new(EC2Instance)
                instance.LaunchTime=inst.LaunchTime
                instance.InstanceId=inst.InstanceId
                instance.InstanceType=inst.InstanceType
                instance.KeyName=inst.KeyName
                instance.State=inst.State.Name
                instance.PublicDns=inst.PublicDnsName
                collection.Add(instance)

	}
	collection.Print()
}
