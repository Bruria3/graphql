package model

type Order struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name"`
	CreatedAt string `bson:"createdAt"`
	Status    string `bson:"status"`
	Quantity  int32  `bson:"quantity"`
}

type UpdateOrderInput struct {
	Name      *string `json:"name,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Status    *string `json:"status,omitempty"`
	Quantity  *int32  `json:"quantity,omitempty"`
}
