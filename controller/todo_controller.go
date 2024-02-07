package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TodoController เป็น struct ที่ใช้ในการควบคุมการทำงานของ API สำหรับ Todo
type TodoController struct{}

// Index เป็น handler function สำหรับ endpoint GET /todos
func (ctrl *TodoController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from TodoController Index!",
	})
}

// Create เป็น handler function สำหรับ endpoint POST /todos
func (ctrl *TodoController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from TodoController Create!",
	})
}

// เพิ่ม handler function อื่นๆ ตามความต้องการ
