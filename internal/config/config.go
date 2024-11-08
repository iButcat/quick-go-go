package config

import "os"

type GoogleCloudServiceAccount struct {
	Type                    string
	ProjectID               string
	PrivateKeyID            string
	PrivateKey              string
	ClientEmail             string
	ClientID                string
	AuthURI                 string
	TokenURI                string
	AuthProviderX509CertURL string
	ClientX509CertURL       string
	UniverseDomain          string
}

type Config struct {
	Port                      string
	MongoURI                  string
	MongoDBName               string
	GoogleCloudServiceAccount GoogleCloudServiceAccount
}

func setEnvOrFallback(envVar string, fallback string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		return value
	}
	return fallback
}

func NewConfig() *Config {
	return &Config{
		Port:        setEnvOrFallback("PORT", "8080"),
		MongoURI:    setEnvOrFallback("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBName: setEnvOrFallback("MONGODB_NAME", "quickgogo"),
		GoogleCloudServiceAccount: GoogleCloudServiceAccount{
			Type:                    setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_TYPE", ""),
			ProjectID:               setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_PROJECT_ID", ""),
			PrivateKeyID:            setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_PRIVATE_KEY_ID", ""),
			PrivateKey:              setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_PRIVATE_KEY", ""),
			ClientEmail:             setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_CLIENT_EMAIL", ""),
			ClientID:                setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_CLIENT_ID", ""),
			AuthURI:                 setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_AUTH_URI", ""),
			TokenURI:                setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_TOKEN_URI", ""),
			AuthProviderX509CertURL: setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_AUTH_PROVIDER_X509_CERT_URL", ""),
			ClientX509CertURL:       setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_CLIENT_X509_CERT_URL", ""),
			UniverseDomain:          setEnvOrFallback("GOOGLE_CLOUD_SERVICE_ACCOUNT_UNIVERSE_DOMAIN", ""),
		},
	}
}
