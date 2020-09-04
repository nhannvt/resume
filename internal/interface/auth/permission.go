package auth

// Permission corresponds to one handler.
type Permission struct {
	handler string
}

// NewPermission returns the pointer of created permission instance.
func NewPermission(handler string) *Permission {
	return &Permission{handler}
}

// HandlerName returns the name of permitted handler.
func (p *Permission) HandlerName() string {
	return p.handler
}

// PermissionList: Array of Permission
type PermissionList []*Permission
