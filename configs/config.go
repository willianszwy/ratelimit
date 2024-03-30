package configs

import "github.com/spf13/viper"

var cfg *Config

type Config struct {
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	IPMaxRequests    int    `mapstructure:"IP_MAX_REQUESTS"`
	IPBlockedTime    int    `mapstructure:"IP_BLOCKED_TIME"`
	TokenName        string `mapstructure:"TOKEN_NAME"`
	TokenMaxRequests int    `mapstructure:"TOKEN_MAX_REQUESTS"`
	TokenBlockedTime int    `mapstructure:"TOKEN_BLOCKED_TIME"`
}

func LoadConfig(path string) (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	viper.SetConfigName("configs")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
