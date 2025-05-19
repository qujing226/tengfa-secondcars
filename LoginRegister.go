package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) {

	userName := c.PostForm("userName")
	password := c.PostForm("password")

	// 使用参数化查询防止SQL注入
	db := DBConnect("tengfadata")
	defer db.Close()
	sql, err := db.Prepare("INSERT INTO user (name, password) VALUES (?, ?)")
	if err != nil {
		c.JSON(500, gin.H{"error": "数据库操作失败"})
		return
	}
	defer sql.Close()

	_, err = sql.Exec(userName, password)
	if err != nil {
		c.JSON(500, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(200, gin.H{"success": true})

}

func Login(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	db := DBConnect("tengfadata")
	defer db.Close()

	sql, err := db.Prepare("select * from user(name,password) values (?,?)")
	if err != nil {
		c.JSON(500, gin.H{"error": "数据库操作失败"})
		log.Fatal(err)
	}
	defer sql.Close()
	_, err = sql.Exec(userName, password)

	if err != nil {
		c.JSON(500, gin.H{"error": "用户名或密码错误"})
	}

}
