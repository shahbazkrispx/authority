package authority

// The link between the roles and permissions
type RolePermission struct {
	ID           uint `json:"id"`            // Unique id (it gets set automatically by the database)
	RoleID       uint `json:"role_id"`       // Role id
	PermissionID uint `json:"permission_id"` // Permission id
}

// TableName sets the table name
func (r RolePermission) TableName() string {
	return auth.TablesPrefix + "role_permissions"
}
