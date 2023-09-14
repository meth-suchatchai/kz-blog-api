package blogmodels

type Seo struct {
	MetaTitle       string `json:"meta_title" form:"meta_title"`
	MetaDescription string `json:"meta_description" form:"meta_description"`
}
