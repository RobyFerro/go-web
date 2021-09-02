package validation

type NewUser struct {
	Name           string `json:"name" valid:"required,alpha"`
	Surname        string `json:"surname" valid:"required,alpha"`
	Username       string `json:"username" valid:"required,alpha"`
	Password       string `json:"password" valid:"required,alpha"`
	RepeatPassword string `json:"repeat-password" valid:"required,alpha"`
}
