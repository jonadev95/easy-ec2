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

func cutAt(count int, targetString *string)string{
	return (*targetString)[:count]
}
