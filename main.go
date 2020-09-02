// File:    main
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/30 21:00
// DESC:    main function

package main

import (
	"GoBlog/model"
	"GoBlog/routes"
)

func main() {
	model.InitMysqlDB()
	routes.InitRouter()
}
