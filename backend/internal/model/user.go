package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"` // ชื่อผู้ใช้
	Password string `json:"password"`                    // รหัสผ่าน
	Name     string `json:"name"`                        // ชื่อเต็ม
	Role     string `json:"role"`                        // บทบาท (เช่น "admin", "user")
}
