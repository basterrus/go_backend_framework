package roles

type Role struct {
	Id          int    `json:"id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description,omitempty"`
}

type CreateRoleDTO struct {
	RoleName    string `json:"role_name"`
	Description string `json:"description,omitempty"`
}

type UpdateRoleDTO struct {
	RoleName    string `json:"role_name"`
	Description string `json:"description,omitempty"`
}

func NewRole(dto CreateRoleDTO) Role {
	return Role{
		RoleName:    dto.RoleName,
		Description: dto.Description,
	}
}

func UpdatedRole(dto UpdateRoleDTO) Role {
	return Role{
		RoleName:    dto.RoleName,
		Description: dto.Description,
	}
}
