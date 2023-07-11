package dto

type LoginReq struct {
	User   string `form:"user" json:"user"`
	Passwd string `form:"passwd" json:"passwd"`
}
type LoginResp struct {
}

type ReSetPasswdReq struct {
	OldPasswd string `form:"old_passwd" json:"old_passwd"`
	NewPasswd string `form:"new_passwd" json:"new_passwd"`
}
type ReSetPasswdResp struct {
}
