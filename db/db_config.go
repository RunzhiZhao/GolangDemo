package db

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func MyDbConfig() DbConfig {
	return DbConfig{
		Username: "root",
		Password: "aa123456",
		Host:     "localhost",
		Port:     "3306",
		DbName:   "order_schema",
	}
}
