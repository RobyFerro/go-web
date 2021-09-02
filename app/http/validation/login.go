package validation

type Credentials struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}
