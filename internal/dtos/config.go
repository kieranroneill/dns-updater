package dtos

type Config struct {
	Auth   AuthConfig   `yaml:"auth"`
	Domain string       `yaml:"domain"`
	Record RecordConfig `yaml:"record"`
}
