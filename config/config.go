package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Env         string `env:"APP_ENV,default=development"`
	AppName     string `env:"APP_NAME,default=starter-api"`
	AppVersion  string `env:"APP_VERSION,default=1.0.0"`
	JobPassword string `env:"JOB_PASSWORD"`
	Redis       Redis
	Port        Port
	JWTConfig   JWTConfig
	Google      Google
	Gotenberg   Gotenberg
	AWS         AWS
	Postgres    Postgres
}

// Redis holds configuration for the Redis.
type Redis struct {
	Address  string `env:"REDIS_ADDRESS"`
	Password string `env:"REDIS_PASSWORD"`
}

// Port holds configuration for project's port.
type Port struct {
	APP string `env:"PORT,default=8080"`
}

// JWTConfig holds configuration for jwt.
type JWTConfig struct {
	Public      string `env:"JWT_PUBLIC,required"`
	Private     string `env:"JWT_PRIVATE,required"`
	Issuer      string `env:"JWT_ISSUER,required"`
	IssuerAdmin string `env:"JWT_ISSUER_ADMIN,required"`
}

// Google holds configuration for the Google.
type Google struct {
	ProjectID         string `env:"GOOGLE_PROJECT_ID"`
	StorageBucketName string `env:"GOOGLE_STORAGE_BUCKET_NAME"`
	StorageEndpoint   string `env:"GOOGLE_STORAGE_ENDPOINT"`
	MapsAPI           string `env:"GOOGLE_MAPS_API"`
}

type Gotenberg struct {
	Host string `env:"GOTENBERG_HOST"`
}

// AWS holds configuration for the AWS.
type AWS struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Region          string `env:"AWS_REGION"`
	BucketName      string `env:"AWS_BUCKET_NAME"`
}

// Postgres holds all configuration for PostgreSQL.
type Postgres struct {
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	MaxOpenConns    string `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN_LIFETIME,default=10m"`
	MaxIdleLifetime string `env:"POSTGRES_MAX_IDLE_LIFETIME,default=5m"`
}

// LoadConfig initiate load config either from env file or os env
func LoadConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
