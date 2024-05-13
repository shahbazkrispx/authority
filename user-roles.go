package authority

// The link between the users and roles
type UserRole struct {
	ID     uint   `json:"id"`      // Unique id (it gets set automatically by the database)
	UserID string `json:"user_id"` // The user id
	RoleID uint   `json:"role_id"` // The role id
}

// TableName sets the table name
func (u UserRole) TableName() string {
	return auth.TablesPrefix + "user_roles"
}
