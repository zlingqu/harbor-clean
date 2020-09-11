package main

import (
	"log"

	"github.com/zlingqu/harbor-clean/cmd"
)

// var (
// 	url         string
// 	user        string
// 	password    string
// 	projectName string
// 	keepNum     int
// 	help        bool
// )

// func init() {
// 	flag.BoolVar(&help, "h", false, "help message")
// 	flag.StringVar(&url, "url", "", "harbor地址，例如https://harbor.abc.com")
// 	flag.StringVar(&user, "user", "", "harbor账号")
// 	flag.StringVar(&password, "password", "", "harbor密码")
// 	flag.StringVar(&projectName, "projectName", "", "项目名。注意：all 表示全部")
// 	flag.IntVar(&keepNum, "keepNum", 0, "每个repo保留的tag个数")
// }

func main() {
	newClean := cmd.NewHarborCleanCommand()
	if err := newClean.Execute(); err != nil {
		log.Fatal(err)
	}

}
