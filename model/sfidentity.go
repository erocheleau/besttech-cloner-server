package model

type Sfidentity struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	Sandbox  bool   `json:"sandbox,omitempty"`
}
