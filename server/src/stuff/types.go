package stuff

type Document struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Mineatura string `json:"mineatura"`
	Body      string `json:"bodyOfDocument"`
}
type Publications struct {
	Size         int        `json:"Size"`
	Publications []Document `json:"Publications"`
	Quantity     int        `json:"Quantity"`
}

type Comment struct {
	
	Username string `json:"username"`
	Body     string `json:"body-comment"`
}
type Comments struct{
	Size int `json:"size"`
	Comments []Comment `json:"comments"`
	Quantity int `json:"Quantity-comments"`
}

type ErrorCode struct{
	Code int
	Message string
}