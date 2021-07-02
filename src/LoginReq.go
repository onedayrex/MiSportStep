package src

//login Object

type LoginReq struct {
	ClientId string `json:"client_id"`
	Password string `json:"password"`
	RedirectUri string `json:"redirect_uri"`
	Token string `json:"token"`
}

func New(password string)*LoginReq  {
	return &LoginReq{
		ClientId:    "HuaMi",
		Password:    password,
		RedirectUri: "https://s3-us-west-2.amazonaws.com/hm-registration/successsignin.html",
		Token:       "access",
	}
}