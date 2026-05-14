package configs

type Secrets struct {
	MySQL MySQL `mapstructure:",squash"`
	// Redis Redis `mapstructure:",squash"`
}

type MySQL struct {
	Host     string `mapstructure:"mysql_host"`
	Port     int    `mapstructure:"mysql_port"`
	User     string `mapstructure:"mysql_user"`
	Password string `mapstructure:"mysql_password"`
	Database string `mapstructure:"mysql_db"`
}

// type Redis struct {
// 	Host     string `mapstructure:"redis_host"`
// 	Port     int    `mapstructure:"redis_port"`
// 	Password string `mapstructure:"redis_password"`
// }
