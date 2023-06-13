package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键ID"  form:"id"`                 // 主键ID
	CreatedAt time.Time      `json:"createAt" gorm:"type:datetime(0);autoCreateTime;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updateAt" gorm:"type:datetime(0);autoUpdateTime;comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                  // 删除时间
}
