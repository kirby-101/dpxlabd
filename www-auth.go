package dpxlabd

type Role uint

const (
	RoleAnyone Role = iota
	RoleAdmin
	RoleUser
	RoleCustomer
)
