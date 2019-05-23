package main
import  (
	"os"
	"fmt"
	"time"
)

func main() {
	test()
}

func test() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs", time.Since(start).Seconds())
}
