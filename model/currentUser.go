package model

type CurrentUser struct {
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Userid      string     `json:"userid"`
	Email       string     `json:"email"`
	Signature   string     `json:"signature"`
	Title       string     `json:"title"`
	Group       string     `json:"group"`
	Tags        []Tags     `json:"tags"`
	NotifyCount int        `json:"notifyCount"`
	UnreadCount int        `json:"unreadCount"`
	Country     string     `json:"country"`
	Access      string     `json:"access"`
	Geographic  Geographic `json:"geographic"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
}
type Tags struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
type Province struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
type City struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
type Geographic struct {
	Province Province `json:"province"`
	City     City     `json:"city"`
}
