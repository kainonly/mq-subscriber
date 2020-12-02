package options

type SubscriberOption struct {
	Identity string `yaml:"identity"`
	Queue    string `yaml:"queue"`
	Url      string `yaml:"url"`
	Secret   string `yaml:"secret"`
}
