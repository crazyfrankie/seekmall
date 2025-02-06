package domain

type Product struct {
	Id          int
	Name        string
	Description string
	Picture     string
	Price       float32
	Uid         int
	Categories  []string
}
