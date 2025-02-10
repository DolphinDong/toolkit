package config

var (
	Cfg Config
)

type Config struct {
	Username string
	Password string
	Host     string
	DB       string
	Other    Other
}
type Other struct {
	Args1  string
	Args2  string
	Args3  string
	Other2 Other2
}

type Other2 struct {
	Args1 string
	Args2 string
	Args3 string
}
