package domain

type Order struct {
	OrderUid		string		`json:"order_uid"`
	TrackNumber		string		`json:"track_number"`
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
	ChrtId			int 	`json:"chrt_id"`
	TrackNumber		string	`json:"track_number"`
	Price			int		`json:"price"`
	Rid				string	`json:"rid"`

}

