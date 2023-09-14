package rpmodels

type RolePermission struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type RolePermissionRequest struct {
}

type RolePermissionResponse struct {
	RolePermissions []RolePermission `json:"role_permissions"`
}
