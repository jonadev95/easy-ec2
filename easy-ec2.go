package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/jonadev95/easy-ec2/instances"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"fmt"
	"os"
	"github.com/jonadev95/easy-ec2/images"
)

var(
	app = kingpin.New("easy-ec2","An easy client for the Amazon EC2")

	json = app.Command("json","")
	ls = app.Command("ls","List all instances")

	rm = app.Command("rm","Remove instances(s)")
	rm_instanceIds = rm.Arg("instance-id","Ids of the instance which should be terminated").Required().Strings()

	listImages = app.Command("images","List Available Images")
	inspect = app.Command("inspect","Show details of instance(s)")

	run = app.Command("run", "Run new Instance(s)")
	run_image	= run.Flag("image","An Alias of an Image (provided by the image command) or the Amazon AMI ID").Required().Short('i').String()
	run_count	= run.Flag("count","Specify How Many Instances should be created").Default("1").Short('c').Int()
	run_keypair	= run.Flag("key-pair","Specify which key pair should be used").Required().Short('k').String()
	run_instanceType= run.Flag("instance-type","Specify Which Instance Type should be used").Short('t').Default("t2.micro").String()

	win_passwd = app.Command("password","Get the password of windows instance(s)")
	win_passwd_ids = win_passwd.Arg("instance-id","Ids of the instance where to get the password").Required().Strings()
	win_passwd_pemFile = win_passwd.Flag("pem-file","Pem File").Short('P').Required().ExistingFile()

	start = app.Command("start", "Start instance(s)")
	start_instanceIds= start.Arg("instance-id","Ids of the instance which should be started").Required().Strings()

	stop = app.Command("stop","Stop instance(s)")
	stop_instanceIds = stop.Arg("instance-id","Ids of the instance which should be started").Required().Strings()

	search = app.Command("search","Search for Images")
	search_filter_name = search.Arg("name","Name of Images").Strings()
	search_filter_image_id = search.Flag("image-id","ID of an Image").Strings()
	search_filter_desc = search.Flag("description","Description of an Image").Strings()
	search_filter_platform = search.Flag("platform","Platform of an Image").Strings()

	configure = app.Command("configure","Set credentials and availability zone")


)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})
	switch kingpin.MustParse(app.Parse(os.Args[1:])){

		case ls.FullCommand():
			instances.Ls(svc)

		case rm.FullCommand():
			Ids := make([]*string, 0)
			for _, instanceId := range *rm_instanceIds {
				Ids = append(Ids, aws.String(instanceId))
			}
			instances.Rm(svc, Ids)

		case run.FullCommand():
			imageId := images.GetImageId(run_image)
			if len(*imageId)==0{
				imageId=run_image
			}
			fmt.Printf("Image ID: %s Count: %d KeyPair: %s instanceType: %s\n",*imageId, *run_count, *run_keypair, *run_instanceType)
			instances.Run(svc, imageId, *run_count, run_keypair, run_instanceType)

		case start.FullCommand():
			Ids := make([]*string,0)
			for _, instanceId := range *start_instanceIds {
				Ids = append(Ids, aws.String(instanceId))
			}
			instances.Start(svc, Ids)

		case stop.FullCommand():
			Ids := make([]*string,0)
			for _, instanceId :=  range *stop_instanceIds {
				Ids = append(Ids, aws.String(instanceId))
			}
			instances.Stop(svc, Ids)
		case listImages.FullCommand():
			images.Print()
		case json.FullCommand():
			images.Json()
		case win_passwd.FullCommand():
			instances.GetPasswd(svc, win_passwd_ids,win_passwd_pemFile)
		case search.FullCommand():
			images.Search(svc, search_filter_name, search_filter_image_id, search_filter_desc, search_filter_platform)
	}
}
