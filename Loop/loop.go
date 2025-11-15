package loop

import (
	"fmt"
	"strings"
)

func Repeated(charachter string, repeatedCount int) string {
	var repeated strings.Builder
	for i := 1; i <= repeatedCount; i++ {
		repeated.WriteString(charachter)
	}
	return repeated.String()
}

func main() {
	fmt.Println(Repeated("a", 10))
}
