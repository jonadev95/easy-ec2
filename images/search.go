package images

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/jonadev95/easy-ec2/formatting"
	"fmt"
)

func Search(svc *ec2.EC2, name *[]string, id *[]string, desc *[]string, platform *[]string){
	params := &ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{},
	}

	if len(*name) != 0{
		var names []*string
		for _, n := range *name{
			names = append(names, aws.String(n))
		}
		filterName:="name"
		filter:=ec2.Filter{
			Name:&filterName,
			Values:names,
		}
		params.Filters = append(params.Filters, &filter)
	}

	if len(*id) != 0{
		var values []*string
		for _, v := range *id{
			values = append(values, aws.String(v))
		}
		filterName:="image-id"
		filter:=ec2.Filter{
			Name:&filterName,
			Values:values,
		}
		params.Filters = append(params.Filters, &filter)
	}

        if len(*desc) != 0{
                var values []*string
                for _, v := range *desc{
                        values = append(values, aws.String(v))
                }
                filterName:="description"
                filter:=ec2.Filter{
                        Name:&filterName,
                        Values:values,
                }
                params.Filters = append(params.Filters, &filter)
        }

        if len(*platform) != 0{
                var values []*string
                for _, v := range *platform{
                        values = append(values, aws.String(v))
                }
                filterName:="platform"
                filter:=ec2.Filter{
                        Name:&filterName,
                        Values:values,
                }
                params.Filters = append(params.Filters, &filter)
        }

	if len(params.Filters) == 0{
		fmt.Println("Please provide a filter")
		return
	}


	resp, err := svc.DescribeImages(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("|%-20s|%-60s|%-60s|\n","Amazon Image ID","Name","Description")
	dash := "-"
	twenty := formatting.GetMultipleStrings(20,&dash)
	sixty := formatting.GetMultipleStrings(60,&dash)
        fmt.Printf("|%-20s|%-60s|%-60s|\n",twenty,sixty,sixty)

	for _, img := range resp.Images{
		desc := ""
		name := ""
		if img.Description != nil{
			desc=*formatting.CutAt(img.Description,60)
		}
		if img.Name != nil{
			name=*formatting.CutAt(img.Name,60)
		}
		fmt.Printf("|%-20s|%-60s|%-60s|\n",*img.ImageId,name,desc)
	}
}
