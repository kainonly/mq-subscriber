package types

type Config struct {
	Debug   string        `yaml:"debug"`
	Listen  string        `yaml:"listen"`
	Mq      MqOption      `yaml:"mq"`
	Logging LoggingOption `yaml:"logging"`
}
