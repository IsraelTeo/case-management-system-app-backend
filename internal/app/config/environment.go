package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Container contiene las variables de entorno para la aplicación en general.
type Container struct {
	App  *App
	DB   *DB
	HTTP *HTTP
	JWT  *JWT
}

// App contiene todas las variables de entorno para la aplicación
type App struct {
	Name string
	Env  string
}

// DB contiene todas las variables de entorno para la base de datos
type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	MinConn  int
	MaxConn  int
}

// HTTP contiene todas las variables de entorno para el servidor HTTP
type HTTP struct {
	Env            string
	URL            string
	Port           string
	AllowedOrigins string
}

// JWT contiene todas las variables de entorno para json web tokens
type JWT struct {
	ExpirationTime int64
	SecretKey      string
}

// New crea una nueva instancia de contenedor
func New() (*Container, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		MinConn:  getEnvAsInt("DB_MIN_CONN", 3),
		MaxConn:  getEnvAsInt("DB_MAX_CONN", 100),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("SERVER_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	jwt := &JWT{
		ExpirationTime: getEnvExpTimeJWTAsInt("EXP_JWT", 10, 64),
		SecretKey:      os.Getenv("API_SECRET"),
	}

	return &Container{
		App:  app,
		DB:   db,
		HTTP: http,
		JWT:  jwt,
	}, nil
}

// getEnvAsInt obtiene el string desde una variable de entorno
// y lo devuelve como int, por defecto retorna un int
func getEnvAsInt(env string, defaultVal int) int {
	valStr := os.Getenv(env)
	if valStr == "" {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		fmt.Println("Error parsing env var", env, "as int:", err)
		return defaultVal
	}

	return val
}

// getEnvExpTimeJwtAsInt obteiene el string desde la variable de entorno "JWT_EXP"
// y lo retorna como int64
func getEnvExpTimeJWTAsInt(env string, base, bitSize int) int64 {
	var defaultVal int64 = 86400

	valStr := os.Getenv(env)
	if valStr == "" {
		return defaultVal
	}

	exp, err := strconv.ParseInt(valStr, base, bitSize)
	if err != nil {
		slog.Error("Error converting JWT_EXP, using default value",
			slog.String("error", err.Error()),
			slog.Int64("default", defaultVal),
		)

		return defaultVal
	}

	return exp
}
