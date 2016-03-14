package instances

import(
	"fmt"
//	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jonadev95/easy-ec2/formatting"
	"time"
)


type EC2Instance struct{
	LaunchTime *time.Time
	InstanceId *string
	InstanceType *string
	KeyName *string
	State *string
	PublicDns *string
}

type EC2InstanceCollection struct{
	Instances []*EC2Instance
}

func NewCollection() *EC2InstanceCollection{
	return &EC2InstanceCollection{}
}

func (collection *EC2InstanceCollection) Add (instance *EC2Instance){
	collection.Instances = append(collection.Instances, instance)
}


func (collection *EC2InstanceCollection) Print(){
	fmt.Printf("|%-20s|%-20s|%-20s|%-20s|%-20s|%-50s|\n","Instance ID","Instance Type","Key Name","Launch Time","State","Public DNS")
        dash := "-"
        dashes := formatting.GetMultipleStrings(20, &dash)
        moreDashes := formatting.GetMultipleStrings(50,&dash)
        fmt.Printf("|%s|%s|%s|%s|%s|%s|\n",dashes,dashes,dashes,dashes,dashes,moreDashes)
        for _, inst :=  range collection.Instances{
		keyName := ""
		if inst.KeyName != nil {
			keyName = *inst.KeyName
		}
		t:= fmt.Sprintf("%d:%d:%d %d-%d-%d", inst.LaunchTime.Hour(),inst.LaunchTime.Minute(), inst.LaunchTime.Second(), inst.LaunchTime.Day(), inst.LaunchTime.Month(), inst.LaunchTime.Year())
		fmt.Printf("|%-20s|%-20s|%-20s|%-20s|%-20s|%-50s|\n",*inst.InstanceId, *inst.InstanceType, keyName,t,*inst.State,*inst.PublicDns)
        }
}
