package blogmodels

type BlogUpdate struct {
	Name             *string   `json:"name" form:"name"`
	Content          *string   `json:"content" form:"content"`
	Category         *Category `json:"category" form:"category"`
	Tags             []Tag     `json:"tags" form:"tags[]"`
	ImageURL         *string   `json:"image_url" form:"image_url"`
	Slug             *string   `json:"slug" form:"slug"`
	Seo              *Seo      `json:"seo" form:"seo"`
	ShortDescription *string   `json:"short_description" form:"short_description"`
}
