package timeline

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// func main() {
// 	// fmt.Println(getUserInput())
// 	x := getUserInput()

// 	fmt.Println("The value is", x)
// }

func DisplayTime() {
	var timeNow time.Time = time.Now()
	fmt.Println(timeNow)
	// fmt.Println("The time right now is:\n", breakDownTime(timeNow))
}

func FloorAndCeil(x float64) (xFloored float64, xCeil float64) {
	xCeil = math.Ceil(x)
	xFloored = math.Floor(x)

	// y = strings.Format("%i (floored): %d\n%i (ceil): %d", x, x_floor, x, x_ceil)
	// y = strconv.ParseString(xFloored) + " (floored)\n" + strconv.ParseString(xCeil) + " (ceil)"
	// y = "Nevermore"
	return
}

func BreakDownTime(x time.Time) (x_ int) {
	// var year int = x.Year()
	// var month int = x.Month()
	// var day int = x.Day()

	// var hour int = x.Hour()
	// var minute int = x.Minute()
	// var second int = x.Second()

	// fmt.Println(year)

	x_ = 1
	return x_
}

func GetUserInput() float64 {
	reader := bufio.NewReader(os.Stdin)

	input, err1 := reader.ReadString('\n')

	if err1 != nil {
		log.Fatal(err1)
	}

	// var val float64

	val, _ := strconv.ParseFloat(input, 64)

	fmt.Println("I got", val)

	return val
}
