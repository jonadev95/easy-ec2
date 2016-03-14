package instances

import(
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jonadev95/easy-ec2/formatting"

)

func Start(svc *ec2.EC2, instanceIDs []*string){
	params := &ec2.StartInstancesInput{
		InstanceIds: instanceIDs,
	}
	resp, err := svc.StartInstances(params)

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("|%-20s|%-20s|%-20s|\n","Instance ID","Current State","Previous State")
	dash := "-"
	dashes := formatting.GetMultipleStrings(20,&dash)
	fmt.Printf("|%s|%s|%s|\n",dashes, dashes, dashes)

	for _,instance := range resp.StartingInstances {
		fmt.Printf("|%-20s|%-20s|%-20s|\n", *instance.InstanceId, *instance.CurrentState.Name, *instance.PreviousState.Name)
	}
}
