package clientmodels

import blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"

type ListBlogResponse struct {
	Blogs []blogmodels.Blog `json:"blogs"`
}

type GetBlogResponse struct {
	blogmodels.Blog `json:"blog"`
}
