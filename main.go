package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// ตัวอย่างการสร้าง API endpoint เพื่อดึงข้อมูล Todo list ทั้งหมด
	router.GET("/todos", func(c *gin.Context) {
		// ทำการเชื่อมต่อกับ MySQL database และดึงข้อมูล Todo list จากนั้นส่งกลับไปยัง client
		c.JSON(http.StatusOK, gin.H{
			"message": "Get all todos",
		})
	})

	// ตัวอย่างการสร้าง API endpoint เพื่อเพิ่ม Todo ใหม่
	router.POST("/todos", func(c *gin.Context) {
		// ทำการเชื่อมต่อกับ MySQL database และเพิ่ม Todo ใหม่ลงในฐานข้อมูล
		c.JSON(http.StatusOK, gin.H{
			"message": "Add new todo",
		})
	})

	router.Run(":8080")
}
