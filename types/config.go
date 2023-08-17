package types

type Config struct {
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBAddress    string `mapstructure:"DB_ADDRESS"`
	DBDriverName string `mapstructure:"DB_DRIVER_NAME"`
}
