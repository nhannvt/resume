package auth

// Role is the collection of a permission.
// It indicates the role which can do some actions for a certain purpose.
type Role struct {
	// name means a role name
	name string

	// A list of permitted action this role can handle.
	permissionList PermissionList
}

func NewRole(name string, permissionList PermissionList) *Role {
	return &Role{name, permissionList}
}

// Name returns a role name
func (r *Role) Name() string {
	return r.name
}

//func (r *Role) PermissionList() PermissionList {
//	return r.permissionList
//}

// IsPermitted checks if given handler is permitted to execute.
func (r *Role) IsPermitted(handler string) bool {
	for _, permission := range r.permissionList {
		if handler == permission.HandlerName() {
			return true
		}
	}

	return false
}

// NewAdminRole creates a role for an administrator.
// This role can do any action for sforum API.
// It will be used by developers, not for clients.
func NewAdminRole() *Role {
	return &Role{
		name:           "admin",
		permissionList: PermissionList{},
	}
}

// NewEditorRole creates a role for tag editor.
// This role can not only read but also write tags.
// It will be used by tag management tool client.
func NewEditorRole() *Role {
	return &Role{
		name:           "editor",
		permissionList: PermissionList{},
	}
}

// NewViewerRole creates a role has only read permisstions.
// It will be used by tag viewer.
func NewViewerRole() *Role {
	return &Role{
		name:           "viewer",
		permissionList: PermissionList{},
	}
}

// NewMonitorRole creates a role for monitoring.
// This role can only permission to read metrics for monitoring.
func NewMonitorRole() *Role {
	return &Role{
		name: "monitor",
	}
}
