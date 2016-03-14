package images

import(
	"fmt"
	"github.com/jonadev95/easy-ec2/formatting"
	"strings"
)

type Image struct{
	Alias string
	Description string
	AMI_Id string
}

type ImageCollection struct{
	Images []Image
}

func Print(){
	col := GetImages()
        fmt.Printf("|%-30s|%-15s|%-80s|\n","Alias","AMI ID","Description")
	dash := "-"
	thirty := formatting.GetMultipleStrings(30, &dash)
	fifteen:= formatting.GetMultipleStrings(15, &dash)
	eighty := formatting.GetMultipleStrings(80, &dash)

        fmt.Printf("|%-30s|%-15s|%-80s|\n",thirty,fifteen,eighty)

	for _, inst := range col.Images{
		fmt.Printf("|%-30s|%-15s|%-80s|\n",inst.Alias,inst.AMI_Id,inst.Description)
	}
}

func GetImageId(alias *string) *string{
	col := GetImages()
	for _, inst := range col.Images{
		if strings.EqualFold(*alias,inst.Alias){
			return &inst.AMI_Id
		}
	}
	dummy := ""
	return &dummy
}

func GetImages() *ImageCollection{
	imageCol := new(ImageCollection)
	arr := imageCol.Images
	arr = append(arr, Image{"amazon","Amazon Linux AMI 2015.09.2 (HVM)","ami-e1398992"})
	arr = append(arr, Image{"redhat","Red Hat Enterprise Linux 7.2 (HVM)","ami-8b8c57f8"})
        arr = append(arr, Image{"suse","SUSE Linux Enterprise Server 12 SP1 (HVM)","ami-f4278487"})
        arr = append(arr, Image{"ubuntu","Ubuntu Server 14.04 LTS (HVM)","ami-f95ef58a"})
        arr = append(arr, Image{"windows","Microsoft Windows Server 2012 R2 Base","ami-8519a9f6"})
        arr = append(arr, Image{"windows2012_r2_sql_expr","Microsoft Windows Server 2012 R2 with SQL Server Express","ami-3201b141"})
        arr = append(arr, Image{"windows2012_r2_sql_web","Microsoft Windows Server 2012 R2 with SQL Server Web","ami-6c02b21f"})
        arr = append(arr, Image{"windows2012_r2_sql","Microsoft Windows Server 2012 R2 with SQL Server Standard","ami-6204b411"})
        arr = append(arr, Image{"windows2012","Microsoft Windows Server 2012 Base","ami-e203b391"})
        arr = append(arr, Image{"windows2012_sql_expr","Microsoft Windows Server 2012 with SQL Server Express","ami-8201b1f1"})
        arr = append(arr, Image{"windows2012_sql_web","Microsoft Windows Server 2012 with SQL Server Web","ami-bbbe0dc8"})
        arr = append(arr, Image{"windows2012_sql","Microsoft Windows Server 2012 with SQL Server Standard","ami-8e06b6fd"})
        arr = append(arr, Image{"windows2008_r2","Microsoft Windows Server 2008 R2 Base","ami-7d00b00e"})
        arr = append(arr, Image{"windows2008_r2_sql_expr_iis","Microsoft Windows Server 2008 R2 with SQL Server Express and IIS","ami-0203b371"})
        arr = append(arr, Image{"windows2008_r2_sql_web","Microsoft Windows Server 2008 R2 with SQL Server Web","ami-6107b712"})
        arr = append(arr, Image{"windows2008_r2_sql","Microsoft Windows Server 2008 R2 with SQL Server Standard","ami-7a06b609"})
        arr = append(arr, Image{"windows2008","Microsoft Windows Server 2008 Base (64-bit)","ami-121eae61"})
        arr = append(arr, Image{"windows2008_32b","Microsoft Windows Server 2008 Base (32-bit)","ami-f918a88a"})
        arr = append(arr, Image{"suse_11","SUSE Linux Enterprise Server 11 SP4 (PV)","ami-fa7cdd89"})
        arr = append(arr, Image{"ubuntu_pv","Ubuntu Server 14.04 LTS (PV)","ami-be5cf7cd"})
        arr = append(arr, Image{"amazon_pv","Amazon Linux AMI 2015.09.2 (PV)","ami-a93484da"})
        arr = append(arr, Image{"windows2003_r2","Microsoft Windows Server 2003 R2 Base (64-bit)","ami-832f9ff0"})
        arr = append(arr, Image{"windows2003_r2_32b","Microsoft Windows Server 2003 R2 Base (32-bit)","ami-af2f9fdc"})

	imageCol.Images=arr
	return imageCol
}
