package configs

type Secrets struct {
	JWTSecret         string `mapstructure:"jwt_secret"`
	GoogleClientID    string `mapstructure:"google_client_id"`
	GoogleClientSecret string `mapstructure:"google_client_secret"`
	GoogleRedirectURL string `mapstructure:"google_redirect_url"`
	Postgres          Postgres `mapstructure:",squash"`
	Redis             Redis    `mapstructure:",squash"`
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
