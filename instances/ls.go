package instances

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jonadev95/easy-ec2/formatting"
)

func Ls(svc *ec2.EC2){
	resp, err := svc.DescribeInstances(nil)
	if err != nil{
		panic(err)
	}
	fmt.Println(resp)
	fmt.Printf("|%-20s|%-20s|%-20s|%-20s|%-20s|%-50s|\n","Instance ID","Instance Type","Key Name","Launch Time","State","Public DNS")
	dash := "-"
	dashes := formatting.GetMultipleStrings(20, &dash)
	moreDashes := formatting.GetMultipleStrings(50,&dash)
	fmt.Printf("|%s|%s|%s|%s|%s|%s|\n",dashes,dashes,dashes,dashes,dashes,moreDashes)
	for idx, _ := range resp.Reservations {
		for _, inst :=  range resp.Reservations[idx].Instances {
			keyName := ""
			if inst.KeyName != nil {
				keyName = *inst.KeyName
			}
			t:= fmt.Sprintf("%d:%d:%d %d-%d-%d",
inst.LaunchTime.Hour(),inst.LaunchTime.Minute(), inst.LaunchTime.Second(), inst.LaunchTime.Day(), inst.LaunchTime.Month(), inst.LaunchTime.Year())
			fmt.Printf("|%-20s|%-20s|%-20s|%-20s|%-20s|%-50s|\n",
*inst.InstanceId, *inst.InstanceType, keyName,t,*inst.State.Name,*inst.PublicDnsName)
		}
	}
}
