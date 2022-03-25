package main

import "Project/todomvc/web"

func main() {
	//需要先建一个数据库，名字:todomvc
	web.Default("databaseName", "password")
}
