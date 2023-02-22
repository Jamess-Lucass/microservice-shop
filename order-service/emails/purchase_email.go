package emails

type PurchaseEmail struct {
	Name    string
	ID      string
	Address string
	Items   []PurchaseEmailItem
}

type PurchaseEmailItem struct {
	Name     string
	Quantity uint
	Price    string
}
