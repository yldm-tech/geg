/**
 * @Author: Evan<suzukaze.hazuki2020@gmail.com>
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/11/30 14:50
 */

package main

import (
	_ "myetc.lantron.ltd/m/docs"
	"myetc.lantron.ltd/m/internal"
)

// @title myetc-api
// @version 1.0.0
// @description This is translator api for lantron ltd.
// @termsOfService https://github.com/lantron ltd-jp
// @contact.name Evan
// @contact.url https://lantron.ltd
// @contact.email suzukaze.hazuki2020@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	application := internal.NewApp()
	defer application.Close()
	application.
		Configure().
		ConnectDB().
		Serve()
}
