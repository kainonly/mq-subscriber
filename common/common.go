package common

type (
	AppOption struct {
		Debug  bool       `yaml:"debug"`
		Listen string     `yaml:"listen"`
		Amqp   AmqpOption `yaml:"amqp"`
	}
	AmqpOption struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Vhost    string `yaml:"vhost"`
	}
	SubscriberOption struct {
		Identity string
		Queue    string
		Url      string
		Secret   string
	}
)
