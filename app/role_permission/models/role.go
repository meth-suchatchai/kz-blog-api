package rpmodels

type Role struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRoleRequest struct {
	Role
}

type CreateRoleResponse struct {
}

type UpdateRoleRequest struct {
	Role
}

type UpdateRoleResponse struct{}
