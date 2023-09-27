package main

import (
	"study_gin/conf"
	"study_gin/dao"
	"study_gin/routes"
)

func main() {
	loading()
	r := routes.NewRouter()
	err := r.Run(conf.HttpPort)
	if err != nil {
		panic(err)
	}
}
func loading() {
	conf.Init()
	dao.MySQLInit()
}
