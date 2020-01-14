package config

type Host struct {
	Disabled bool            `yaml:"disabled"`
	Address string           `yaml:"address"`
}

type AdapterConfig struct {
	Protocol string          `yaml:"protocol"`
	MessageDispatcher Host   `yaml:"grpc"`
	Kafka Host               `yaml:"kafka"`
}

type FrontConfig struct {
	FrontApi Host            `yaml:"api"`
	SessionApi Host          `yaml:"session"`
	Adapter AdapterConfig    `yaml:"adapter"`
}

type MySQL struct {
	Disabled bool            `yaml:"disabled"`
	User string              `yaml:"user"`
	Password string          `yaml:"password"`
	Host string              `yaml:"host"`
	DBName string            `yaml:"dbname"`
	Port int                 `yaml:"port"`
}

type ServiceConfig struct {
	MySQL MySQL              `yaml:"mysql"`    // database
	MongoDB Host             `yaml:"mongodb"`  // history
}

type ApiConfig struct {
	Protocols []string       `yaml:"protocols"`
	Grpc Host                `yaml:"grpc"`
	Http Host                `yaml:"http"`
}

type Config struct {
	Front FrontConfig        `yaml:"front"`
	Service ServiceConfig    `yaml:"service"`
	Api ApiConfig            `yaml:"api"`
}
