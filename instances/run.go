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
			},
		},
	}

	resp,err := svc.RunInstances(params)

	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
