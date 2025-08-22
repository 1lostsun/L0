package model

import "time"

type Order struct {
	ID                uint      `gorm:"primaryKey" json:"-"`
	OrderUID          string    `gorm:"uniqueIndex" json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OOFShard          string    `json:"oof_shard"`

	Delivery Delivery `gorm:"foreignKey:OrderID" json:"delivery"`
	Payment  Payment  `gorm:"foreignKey:OrderID" json:"payment"`
	Items    []Item   `gorm:"foreignKey:OrderID" json:"items"`
}

type Delivery struct {
	ID      uint   `gorm:"primaryKey" json:"-"`
	OrderID uint   `gorm:"index" json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	ID           uint   `gorm:"primaryKey" json:"-"`
	OrderID      uint   `gorm:"index" json:"-"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	OrderID     uint   `gorm:"index" json:"-"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NM_ID       int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
