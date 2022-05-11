package response

type SignUp struct {
	UID string
}

type SignIn struct {
	AccessToken string `json:"accessToken"`
}
