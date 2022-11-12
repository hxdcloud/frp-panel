package model

type LoginParams struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	AutoLogin bool   `form:"autoLogin" json:"autoLogin" binding:"required"`
	Type      string `form:"type" json:"type" binding:"required"`
}

type LoginResult struct {
	Status           string `form:"status" json:"status"`
	Type             string `form:"type" json:"type"`
	CurrentAuthority string `form:"currentAuthority" json:"currentAuthority"`
	Token            string `form:"token" json:"token"`
}
