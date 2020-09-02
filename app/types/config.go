package types

type Config struct {
	Debug   bool          `yaml:"debug"`
	Listen  string        `yaml:"listen"`
	Amqp    string        `yaml:"amqp"`
	Logging LoggingOption `yaml:"logging"`
}
