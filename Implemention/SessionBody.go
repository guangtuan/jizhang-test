package Implemention

type SessionBody struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	UserId   int    `json:"userId"`
}

var sessionBody SessionBody

func SetSession(sessionFromResp SessionBody) {
	sessionBody = sessionFromResp
}

func getSession() SessionBody {
	return sessionBody
}
