package dao

type Consignee struct {
	Phone         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Id        int `gorm:"primaryKey,autoIncrement"`
	UserId    int32
	OrderId   string
	Currency  string      `gorm:"type:varchar(128)"`
	Consignee Consignee   `gorm:"embedded"`
	Items     []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

type OrderItem struct {
	Id           int    `gorm:"primaryKey,autoIncrement"`
	OrderIdRefer string `gorm:"type:varchar(100);index"`
	Quantity     int
	ProductId    int
	Cost         int64 `gorm:"decimal(10, 2)"`
}
