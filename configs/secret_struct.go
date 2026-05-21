package configs

type Secrets struct {
	Redis     Redis  `mapstructure:",squash"`
	JWTSecret string `mapstructure:"jwt_secret"`
	Postgres Postgres `mapstructure:",squash"`
	Redis    Redis    `mapstructure:",squash"`
}

type Postgres struct {
	Host     string `mapstructure:"postgres_host"`
	Port     int    `mapstructure:"postgres_port"`
	User     string `mapstructure:"postgres_user"`
	Password string `mapstructure:"postgres_password"`
	Database string `mapstructure:"postgres_db"`
}

type Redis struct {
	Host     string `mapstructure:"redis_host"`
	Port     int    `mapstructure:"redis_port"`
	Password string `mapstructure:"redis_password"`
	Database int    `mapstructure:"redis_db"`
}
