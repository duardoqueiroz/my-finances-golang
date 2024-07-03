package keyvalues

import "fmt"

const (
	AdminRole = iota
	MemberRole
)

type Role struct {
	value string
}

func NewRole(value int) (*Role, error) {
	switch value {
	case AdminRole:
		return &Role{value: "admin"}, nil
	case MemberRole:
		return &Role{value: "member"}, nil
	default:
		return nil, fmt.Errorf("invalid role")
	}
}

func (r *Role) Value() string {
	return r.value
}
