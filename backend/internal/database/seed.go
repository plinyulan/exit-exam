package database

import (
	"time"

	"github.com/plinyulan/exit-exam/internal/model"
	"gorm.io/gorm"
)

func SeedIfEmpty(db *gorm.DB) error {
	var count int64
	db.Model(&model.Politician{}).Count(&count)
	if count > 0 {
		return nil
	}

	// Users (Authentication แบบง่าย)
	admin := model.User{
		Username: "admin",
		Password: "admin123", // ข้อสอบอนุญาตแบบง่าย
		Name:     "John Doe",
		Role:     "admin",
	}
	user := model.User{
		Username: "user",
		Password: "user123",
		Name:     "Jane Smith",
		Role:     "user",
	}
	if err := db.Create(&admin).Error; err != nil {
		return err
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// Politicians >= 5
	pols := []model.Politician{
		{PoliticianCode: "12345678", Name: "Somchai Prasert", Party: "Future Party"},
		{PoliticianCode: "87654321", Name: "Suda Kanya", Party: "People First"},
		{PoliticianCode: "23456789", Name: "Anan Chai", Party: "Green Nation"},
		{PoliticianCode: "98765432", Name: "Worawat Noon", Party: "Unity Party"},
		{PoliticianCode: "34567890", Name: "Pimchanok Dee", Party: "Progress Alliance"},
	}
	if err := db.Create(&pols).Error; err != nil {
		return err
	}

	// Campaigns
	camps := []model.Campaign{
		{Year: 2566, District: "Bangkok เขต 1"},
		{Year: 2566, District: "Chiang Mai เขต 2"},
		{Year: 2566, District: "Phuket เขต 3"},
		{Year: 2566, District: "Khon Kaen เขต 4"},
		{Year: 2566, District: "Nakhon Ratchasima เขต 5"},
	}
	if err := db.Create(&camps).Error; err != nil {
		return err
	}

	// Promises >= 10, ครบทุกสถานะ
	// statuses: not_started, in_progress, failed
	now := time.Now()
	promises := []model.Promise{
		{PoliticianID: pols[0].ID, CampaignID: camps[0].ID, Detail: "ลดค่าโดยสารสาธารณะ 20%", AnnouncedAt: now.AddDate(0, 0, -30), Status: "not_started"},
		{PoliticianID: pols[0].ID, CampaignID: camps[0].ID, Detail: "เพิ่มพื้นที่สีเขียวในเมือง", AnnouncedAt: now.AddDate(0, 0, -25), Status: "in_progress"},
		{PoliticianID: pols[1].ID, CampaignID: camps[1].ID, Detail: "ปรับปรุงระบบระบายน้ำ", AnnouncedAt: now.AddDate(0, 0, -20), Status: "failed"},
		{PoliticianID: pols[1].ID, CampaignID: camps[1].ID, Detail: "สนับสนุนทุนการศึกษา", AnnouncedAt: now.AddDate(0, 0, -18), Status: "not_started"},
		{PoliticianID: pols[2].ID, CampaignID: camps[2].ID, Detail: "พัฒนาศูนย์สุขภาพชุมชน", AnnouncedAt: now.AddDate(0, 0, -15), Status: "in_progress"},
		{PoliticianID: pols[2].ID, CampaignID: camps[2].ID, Detail: "สร้างทางจักรยานปลอดภัย", AnnouncedAt: now.AddDate(0, 0, -14), Status: "not_started"},
		{PoliticianID: pols[3].ID, CampaignID: camps[3].ID, Detail: "เพิ่ม CCTV จุดเสี่ยง", AnnouncedAt: now.AddDate(0, 0, -12), Status: "in_progress"},
		{PoliticianID: pols[3].ID, CampaignID: camps[3].ID, Detail: "ลดขยะพลาสติกในเขต", AnnouncedAt: now.AddDate(0, 0, -10), Status: "failed"},
		{PoliticianID: pols[4].ID, CampaignID: camps[4].ID, Detail: "ยกระดับทักษะดิจิทัลประชาชน", AnnouncedAt: now.AddDate(0, 0, -9), Status: "not_started"},
		{PoliticianID: pols[4].ID, CampaignID: camps[4].ID, Detail: "ปรับปรุงถนนสายหลัก", AnnouncedAt: now.AddDate(0, 0, -7), Status: "in_progress"},
	}
	if err := db.Create(&promises).Error; err != nil {
		return err
	}

	// Promise Updates (หลายครั้งได้)
	updates := []model.PromiseUpdate{
		{PromiseID: promises[1].ID, Note: "เริ่มปลูกต้นไม้ในสวนสาธารณะ", UpdatedAt: now.AddDate(0, 0, -20)},
		{PromiseID: promises[1].ID, Note: "ขยายพื้นที่ปลูกต้นไม้เพิ่มเติม", UpdatedAt: now.AddDate(0, 0, -10)},
		{PromiseID: promises[4].ID, Note: "จัดซื้ออุปกรณ์ทางการแพทย์ใหม่", UpdatedAt: now.AddDate(0, 0, -12)},
		{PromiseID: promises[4].ID, Note: "เพิ่มจำนวนเจ้าหน้าที่ดูแลสุขภาพ", UpdatedAt: now.AddDate(0, 0, -5)},
		{PromiseID: promises[6].ID, Note: "ติดตั้ง CCTV ในพื้นที่แรก", UpdatedAt: now.AddDate(0, 0, -8)},
		{PromiseID: promises[6].ID, Note: "ขยายการติดตั้ง CCTV ครอบคลุมมากขึ้น", UpdatedAt: now.AddDate(0, 0, -3)},
	}
	return db.Create(&updates).Error
}
