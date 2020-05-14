package configuration

type Configuration struct {
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
}
