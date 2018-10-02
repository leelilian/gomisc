package misc

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now()
	s := ""
	writer := bufio.NewWriter(file)
	for i := 1; i < 10000001; i++ {
		
		s = ":hello world\n"
		
		writer.WriteString(s)
		
	}
	writer.Flush()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
