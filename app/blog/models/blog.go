package blogmodels

type Blog struct {
	Id               int64    `json:"id" form:"id"`
	Name             string   `json:"name" form:"name"`
	Content          string   `json:"content" form:"content"`
	Category         Category `json:"category" form:"category"`
	Tags             []Tag    `json:"tags" form:"tags[]"`
	ImageURL         string   `json:"image_url" form:"image_url"`
	Slug             string   `json:"slug" form:"slug"`
	Seo              `json:"seo" form:"seo"`
	Views            int    `json:"views" form:"views"`
	ShortDescription string `json:"short_description" form:"short_description"`
}

type CreateBlogRequest struct {
	Blog
}

type CreateBlogResponse struct {
}

type GetBlogResponse struct {
	Blog
}

type DeleteBlogRequest struct {
}
type DeleteBlogResponse struct {
}

type ListBlogResponse struct {
	Blogs []Blog `json:"blogs"`
}
