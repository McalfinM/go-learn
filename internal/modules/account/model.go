package account

type Role struct {
	RoleID   int    `json:"id"`
	RoleUuid string `json:"role_uuid"`
	Name     string `json:"name"`
}

type User struct {
	ID       int64  `json:"-"`
	UserUUID string `json:"user_uuid"`

	Email    string `json:"email"`
	Password string `json:"-"`

	RoleID   int64 `json:"-"`
	RoleName string
}
