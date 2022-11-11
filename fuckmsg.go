package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

var logo string = `

___________             __        _____                 
\_   _____/_ __   ____ |  | __   /     \   ______ ____  
 |    __)|  |  \_/ ___\|  |/ /  /  \ /  \ /  ___// ___\ 
 |     \ |  |  /\  \___|    <  /    Y    \\___ \/ /_/  >
 \___  / |____/  \___  >__|_ \ \____|__  /____  >___  / 
     \/              \/     \/         \/     \/_____/  

`

func main() {
	var msg string
	var times int
	fmt.Println(logo)
	fmt.Printf("要輸入的內容:")
	fmt.Scanln(&msg)

	fmt.Printf("次數:")
	fmt.Scanln(&times)

	fmt.Println("5秒內點擊輸入框")

	for waiting := 5; waiting >= 1; waiting-- {
		fmt.Println(waiting)
		time.Sleep(1 * time.Second)
	}

	for i := 0; i <= times-1; i++ {
		robotgo.TypeStr(msg)
		robotgo.KeyTap(`enter`)

	}
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key
}
