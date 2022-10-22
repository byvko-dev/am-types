package users

import "github.com/byvko-dev/am-core/helpers/slices"

type Role string

type Roles []Role

const (
	RoleMasterAdmin Role = "masterAdmin"
	RoleModerator   Role = "moderator"

	RoleDeveloper Role = "developer"

	RolePaidPlus        Role = "amPlus"
	RolePaidSupporter   Role = "amSupporter"
	RolePaidGuildMaster Role = "amGuildMaster"
	RolePaidTheOne      Role = "amTheOne"
)

func (all *Roles) Includes(r Role) bool {
	return slices.Contains(*all, r) > -1
}

func (all *Roles) Add(r Role) {
	if all.Includes(r) {
		return
	}
	*all = append(*all, r)
}

func (all *Roles) Remove(r Role) {
	index := slices.Contains(*all, r)
	if index == -1 {
		return
	}
	*all = append((*all)[:index], (*all)[index+1:]...)
}
