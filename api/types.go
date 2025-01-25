package api

type Profile struct {
	Username string `ini:"username" validate:"required"`
	Password string `ini:"password" validate:"required"`
	Addr     string `ini:"addr"`
	Port     string `ini:"port"`
}

type Node struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Online    bool   `json:"online"`
	Ip        string `json:"ip"`
}
