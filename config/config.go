package config

import "fmt"

type AppMode string

const (
	AppModeDev  AppMode = "development"
	AppModeProd AppMode = "production"
	AppModeTest AppMode = "test"
)

type Config struct {
	AppName string
	AppMode AppMode
	Port    string

	DBType   string
	DBSource string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	JWTSecretKey string
	JWTLifespan  int

	UploadPath string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading.env file")
	}

	return &Config{
		AppName: getEnv("APP_NAME", "OnlineShop"),
		AppMode: AppMode(getEnv("APP_MODE", string(AppModeDev))),
		Port:    getEnv("PORT", "8080"),

		DBType:   getEnv("DB_TYPE", "sqlite"),
		DBSource: getEnv("DB_SOURCE", "./online_shop.db"),

		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		JWTSecretKey: getEnv("JWT_SECRET_KEY", "secret"),
		JWTLifespan:  getEnvInt("JWT_LIFESPAN", 24),

		UploadPath: getEnv("UPLOAD_PATH", "./uploads"),
	}, nil
}
