package web

import (
	"Project/todomvc/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Default(name string, password string) {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/**/*")
	r.Static("/static", "../static")
	r.GET("/todomvc", func(c *gin.Context) {
		//主页
		c.HTML(http.StatusOK, "todomvc", "")
	})
	api(r)
	handle(r, name, password)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func api(r *gin.Engine) {
	api := r.Group("/todomvc")
	api.GET("/addTask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addTask.html", "")
	})
	api.GET("/deleteTask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deleteTask.html", "")
	})
	api.GET("/queryAllTask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "queryTask.html", "")
	})
	api.GET("/selectTask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "selectTask.html", "")
	})
	api.GET("/signTask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signTask.html", "")
	})
}

type JsonData struct {
	NewTask            string `form:"newTask" json:"newTask"`
	DeleteTask         string `form:"deleteTask" json:"deleteTask"`
	QueryTaskCondition string `form:"queryTaskCondition" json:"queryTaskCondition"`
	SignTask           string `form:"signTask" json:"signTask,omitempty"`
	SignTaskCondition  string `form:"signTaskCondition" json:"signTaskCondition"`
	SelectTask         string `form:"selectTask" json:"selectTask"`
	TaskCondition      string `form:"taskCondition" json:"taskCondition"`
}

//处理api
func handle(r *gin.Engine, name string, password string) {
	r.POST("/todomvc/handle", func(c *gin.Context) {
		var jsonData = JsonData{}
		err := c.ShouldBind(&jsonData)
		if err != nil {
			fmt.Println(err)
		}
		c.HTML(http.StatusOK, "handle.html", "")
		//链接数据库
		db, err := database.Client(name, password)
		if err != nil {
			fmt.Println(err)
		}
		//自动迁移
		err = database.AutoMigrate(db)
		if err != nil {
			fmt.Println(err)
		}
		//判断执行
		//添加
		if jsonData.NewTask != "" {
			database.AddTask(db, jsonData.NewTask)
			result := database.Select(db)
			c.JSON(http.StatusOK, result)
			//删除
		} else if jsonData.DeleteTask != "" {
			database.Delete(db, jsonData.DeleteTask)
			result := database.Select(db)
			c.JSON(http.StatusOK, result)
			//查询
		} else if jsonData.QueryTaskCondition != "" {
			l, err := strconv.Atoi(jsonData.QueryTaskCondition)
			if err != nil {
				fmt.Println(err)
			}
			result := database.SelectAll(db, l)
			c.JSON(http.StatusOK, result)
			//标记
		} else if jsonData.SignTask != "" && jsonData.SignTaskCondition != "" {
			l, err := strconv.Atoi(jsonData.SignTaskCondition)
			if err != nil {
				fmt.Println(err)
			}
			database.Sign(db, jsonData.SignTask, l)
			result := database.Select(db)
			c.JSON(http.StatusOK, result)
			//like查询
		} else if jsonData.SelectTask != "" && jsonData.TaskCondition != "" {
			l, err := strconv.Atoi(jsonData.TaskCondition)
			if err != nil {
				fmt.Println(err)
			}
			result := database.SelectLike(db, jsonData.SelectTask, l)
			c.JSON(http.StatusOK, result)
		}
	})

}
