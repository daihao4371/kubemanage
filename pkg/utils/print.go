package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Blue = color.New(color.FgHiBlue, color.Bold).SprintFunc()
)

// PrintLogo 服务启动打印信息
func PrintLogo() {
	fmt.Println(Blue(`
 /$$   /$$           /$$                 /$$      /$$                                                  
| $$  /$$/          | $$                | $$$    /$$$                                                  
| $$ /$$/  /$$   /$$| $$$$$$$   /$$$$$$ | $$$$  /$$$$  /$$$$$$  /$$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$ 
| $$$$$/  | $$  | $$| $$__  $$ /$$__  $$| $$ $$/$$ $$ |____  $$| $$__  $$ |____  $$ /$$__  $$ /$$__  $$
| $$  $$  | $$  | $$| $$  \ $$| $$$$$$$$| $$  $$$| $$  /$$$$$$$| $$  \ $$  /$$$$$$$| $$  \ $$| $$$$$$$$
| $$\  $$ | $$  | $$| $$  | $$| $$_____/| $$\  $ | $$ /$$__  $$| $$  | $$ /$$__  $$| $$  | $$| $$_____/
| $$ \  $$|  $$$$$$/| $$$$$$$/|  $$$$$$$| $$ \/  | $$|  $$$$$$$| $$  | $$|  $$$$$$$|  $$$$$$$|  $$$$$$$
|__/  \__/ \______/ |_______/  \_______/|__/     |__/ \_______/|__/  |__/ \_______/ \____  $$ \_______/
                                                                                    /$$  \ $$          
                                                                                   |  $$$$$$/          
                                                                                    \______/           
`))
}

func Must(err error) {
	panic(err)
}
