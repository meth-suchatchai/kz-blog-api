package rpmodels

type Permission struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type CreatePermissionRequest struct {
	Permission
}

type CreatePermissionResponse struct {
}

type UpdatePermissionRequest struct {
	Permission
}

type UpdatePermissionResponse struct {
}
