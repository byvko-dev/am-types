package users

type Role string

const (
	RoleMasterAdmin Role = "masterAdmin"
	RoleDeveloper   Role = "developer"

	RolePaidPlus        Role = "amPlus"
	RolePaidSupporter   Role = "amSupporter"
	RolePaidGuildMaster Role = "amGuildMaster"
	RolePaidTheOne      Role = "amTheOne"
)
