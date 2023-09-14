package blogmodels

type Category struct {
	Id   int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
