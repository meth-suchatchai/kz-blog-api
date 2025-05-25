package gormdb

import (
	"github.com/meth-suchatchai/kurostatemachine"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
	"gorm.io/gorm"
)

type Client interface {
	Migrate() error
	Seed()
	ORM() *gorm.DB
	/* Blog */
	GetBlogById(id uint) (*dbmodels.Blog, error)
	GetContentBySlug(slug string) (*dbmodels.Blog, error)
	CreateBlog(data *dbmodels.Blog) error
	DeleteBlog(id uint) error
	//UpdateBlog(data *dbmodels.Blog) error
	UpdateBlog(id uint, data map[string]interface{}) error
	ListBlog(opts ...int) (*[]dbmodels.Blog, error)
	ListPopularTag() (*[]dbmodels.Tag, error)
	CreateTag(data *dbmodels.Tag) error
	CreateCategory(data *dbmodels.Category) error
	ListCategory() (*[]dbmodels.Category, error)
	UpdateCategory(data *dbmodels.Category) error
	CountViews(slug string) (int, error)

	/* Role */
	CreateRole(r *dbmodels.Role) error
	CreatePermission(r *dbmodels.Permission) error
	GetRoles(opts ...int) (*[]dbmodels.Role, error)
	GetRolePermission() (*[]dbmodels.Role, error)
	GetPermission(permissionCode string) (*dbmodels.Permission, error)
	AssignRoleToUser(roleId uint, userId uint) (*dbmodels.Role, error)
	AssignPermissionToRole(role string)

	/* User */
	ListUser() ([]dbmodels.User, error)
	GetUser(id uint) (*dbmodels.User, error)
	DeleteUser(id uint) bool
	GetUserByMobileNumber(mobileNumber string) (*dbmodels.User, error)
	CreateUser(data *dbmodels.User) (*dbmodels.User, error)
	UpdateUser(id uint, params map[string]interface{}) error
	UpdateTFAColumn(id uint, provision string, enabled bool) error
	VerifyUser(id uint) error
	GetUserPermission(userId uint, permissionId uint) bool

	GetUserAuthenticationByMobile(mobileNumber, countryCode string) (*dbmodels.UserAuthentication, error)
	GetUserAuthenticationById(id uint, token string) (*dbmodels.UserAuthentication, error)
	CreateUserAuthentication(data *dbmodels.UserAuthentication) (*dbmodels.UserAuthentication, error)
	UpdateUserAuthentication(field *dbmodels.UpdateUserAuthentication) bool
	CreateOrUpdateUserAuthentication(user *dbmodels.User, auth *dbmodels.UpdateUserAuthentication) error

	CreateScene(data *dbmodels.Scene) (*dbmodels.Scene, error)
	UpdateScene(id uint, params map[string]interface{}) error
	UpdateStatusScene(id uint, status kurostatemachine.State) error
}
