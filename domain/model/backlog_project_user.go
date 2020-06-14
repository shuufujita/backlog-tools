package model

// BacklogProjectUser backlog project user
type BacklogProjectUser struct {
	ID          int64  `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int64  `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}
