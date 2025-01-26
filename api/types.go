package api

type Profile struct {
	Username string `ini:"username" validate:"required"`
	Password string `ini:"password" validate:"required"`
}

type Config struct {
	Username string
	Password string
	Addr     string
	Port     string
}

type Node struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	GivenName   string   `json:"given_name"`
	Online      bool     `json:"online"`
	IPAddresses []string `json:"ip_addresses"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt Timestamp `json:"created_at"`
}

type Timestamp struct {
	Seconds int64 `json:"seconds"`
}
