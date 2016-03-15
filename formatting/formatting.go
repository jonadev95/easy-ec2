package formatting

import(
	"bytes"
)

func GetMultipleStrings(count int, targetString *string) string {
	var buffer bytes.Buffer
	for i:=0; i<count; i++ {
		buffer.WriteString(*targetString)
	}
	return buffer.String()
}

func CutAt(targetString *string, count int)*string{
	if len(*targetString)<count{
		return targetString
	}
	ret := (*targetString)[:count]
	return &ret
}
