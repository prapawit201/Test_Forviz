package entity

type Book struct {
	ID          int
	UUID        string
	Name        string
	Author      string
	Type        string
	Price       float64
	Zone        *string
	BorrowCount int
	IsBorrow    bool
	CreatedAt   string
	UpdatedAt   string
}

type BookPaginate struct {
	Page        int
	Limit       int
	FilterBy    *string
	FilterValue *string
	Sort        string
}
