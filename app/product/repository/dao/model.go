package dao

type ProductDraft struct {
	Id          int `gorm:"primaryKey,autoIncrement"`
	Name        string
	Description string
	Picture     string
	Status      string
	Price       float32
	Uid         int
	Ctime       int64
	Utime       int64
	Categories  []Category `gorm:"many2many:product_draft_category"`
}

type ProductLive struct {
	Id          int `gorm:"primaryKey,autoIncrement"`
	Name        string
	Description string
	Picture     string
	Price       float32
	Status      string // 'live', 'offline'
	Stock       int
	Uid         int
	Ctime       int64
	Utime       int64
	Categories  []Category `gorm:"many2many:product_live_category"`
}

type Category struct {
	Id            int `gorm:"primaryKey,autoIncrement"`
	Name          string
	Description   string
	DraftProducts []ProductDraft `gorm:"many2many:product_draft_category"`
	LiveProducts  []ProductLive  `gorm:"many2many:product_live_category"`
}
