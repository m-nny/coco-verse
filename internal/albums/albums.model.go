package albums

type Album struct {
	ID     string  `json:"id" gorm:"primarykey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
