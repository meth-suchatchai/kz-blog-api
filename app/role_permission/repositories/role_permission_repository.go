package rprepositories

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
	"github.com/kuroshibaz/lib/errors"
)

func (repo *defaultRepository) RolePermission() (*[]rpmodels.RolePermission, *fiber.Error) {
	roles, err := repo.db.GetRolePermission()
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	var rolePermissions []rpmodels.RolePermission
	for _, role := range *roles {
		rp := rpmodels.RolePermission{
			Id:   int64(role.ID),
			Name: role.Name,
		}

		ps := make([]rpmodels.Permission, 0)
		for _, permission := range role.Permission {
			m := rpmodels.Permission{
				Id:   int64(permission.ID),
				Name: permission.Name,
				Code: permission.Code,
			}
			ps = append(ps, m)
		}

		rp.Permissions = ps
		rolePermissions = append(rolePermissions, rp)
	}

	return &rolePermissions, nil
}
