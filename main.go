// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	loc, _ := time.LoadLocation("Asia/Seoul")
// 	const longForm = "Jan 2, 2006 at 3:04pm"
// 	t1, _ := time.ParseInLocation(longForm, "Jun 14, 2021 at 10:10pm", loc)
// 	fmt.Println(t1, t1.Location(), t1.UTC())

// 	const shortForm = "2006-Jun-02"
// 	t2, _ := time.Parse(shortForm, "2021-Jun-14")
// 	fmt.Println(t2, t2.Location())

// 	t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(t3, t3.Location())

//		d := t2.Sub(t1)
//		fmt.Println(d)
//	}
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 표준 입력 스트림으로 값을 읽어 올 수 있는 객체를 만든다.
var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		// 표준 입력 스트림을 비워 준다.
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Println("숫자를 입력해 주세요")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력해주세요")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하드립니다. 시도한 횟수 :", cnt)
				break
			}
			cnt++
		}
	}
}

/*
	입력 스트림 활용 .
	및 시간
*/
