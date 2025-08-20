package model

type OrderRequest struct {
	OrderUID        string          `json:"order_uid" validate:"required"`
	TrackNumber     string          `json:"track_number" validate:"required"`
	Entry           string          `json:"entry" validate:"required"`
	Delivery        DeliveryRequest `json:"delivery" validate:"required"`
	Payment         PaymentRequest  `json:"payment" validate:"required"`
	Items           []ItemRequest   `json:"items" validate:"required,dive,gt=0"`
	Locale          string          `json:"locale" validate:"required"`
	CustomerID      string          `json:"customer_id" validate:"required"`
	DeliveryService string          `json:"delivery_service" validate:"required"`
}

type DeliveryRequest struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required,e164"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type PaymentRequest struct {
	Transaction string `json:"transaction" validate:"required"`
	Currency    string `json:"currency" validate:"required"`
	Provider    string `json:"provider" validate:"required"`
	Bank        string `json:"bank" validate:"required"`
}

type ItemRequest struct {
	ChrtID      int    `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required,gt=0"`
	RID         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"min=0"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price" validate:"required,gt=0"`
	NmID        int    `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}
