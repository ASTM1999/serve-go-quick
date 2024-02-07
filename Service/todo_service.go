package service

import (
	"fmt"
)

// TodoService เป็น struct ที่ใช้ในการจัดการธุรกรรมที่เกี่ยวข้องกับ Todo
type TodoService struct{}

// GetAllTodos เป็นฟังก์ชันที่ใช้ในการดึงข้อมูลทั้งหมดของ Todo จากฐานข้อมูล
func (svc *TodoService) GetAllTodos() {
	fmt.Println("Fetching all todos from the database...")
}

// CreateTodo เป็นฟังก์ชันที่ใช้ในการสร้าง Todo ใหม่ลงในฐานข้อมูล
func (svc *TodoService) CreateTodo() {
	fmt.Println("Creating a new todo in the database...")
}
