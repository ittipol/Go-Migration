package repository

type Product struct {
	ID       int
	Name     string `gorm:"size:100"`
	Quantity int
}
