package blogrepositories

import (
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) UpdateBlog(id uint, params *blogmodels.BlogUpdate) errors.Error {
	updates := make(map[string]interface{})

	if params.Name != nil {
		updates["name"] = *params.Name
	}

	if params.Slug != nil {
		updates["slug"] = *params.Slug
	}

	if params.Category != nil {
		updates["category"] = *params.Category
	}

	if params.Content != nil {
		updates["content"] = *params.Content
	}

	if params.Seo != nil {
		updates["seo"] = *params.Seo
	}

	if params.ShortDescription != nil {
		updates["short_description"] = *params.ShortDescription
	}

	err := repo.orm.UpdateBlog(id, updates)
	if err != nil {
		return errors.NewDefaultError(err)
	}

	return nil
}
