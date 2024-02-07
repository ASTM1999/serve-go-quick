package model

// Todo เป็นโครงสร้างข้อมูลที่ใช้ในการแทน Todo ในฐานข้อมูล
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// สามารถเพิ่มโครงสร้างข้อมูลอื่นๆ ตามความต้องการ
