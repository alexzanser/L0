package domain

type Order struct {
	Order_uid		string		`json:"order_uid"`
	Track_number	string		`json:"track_number"`
	Entry			string		`json:"entry"`
	Delivery		Delivery	`json:"delivery"`
	Items			[]Item		`json:"items"`	
}

type Delivery struct {
	Name	string	`json:"name"`
	Phone	string	`json:"phone"`
	Zip		string	`json:"zip"`
	City	string	`json:"city"`
	Address	string	`json:"address"`
	Region	string	`json:"region"`
	Email	string	`json:"email"`
}

type Item struct {
	Chrt_id	int 
}

