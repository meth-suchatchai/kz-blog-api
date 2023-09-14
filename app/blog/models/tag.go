package blogmodels

type Tag struct {
	Id   int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Ord  int    `json:"ord" form:"ord"`
}
