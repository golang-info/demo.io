package main
import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

//declare count variable, which is accessed by all the routine instances
var count = 0

//copies count to temp, do some processing(increment) and store back to count
//random deploy is added between reading and writing of count variable
func process(n int) {
	//loop incrementing the count ty 10
	for i:=0; i<10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

func main() {
	//loop calling the process() 3 times
	for i:=0; i<4; i++ {
		go process(i)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count)
}