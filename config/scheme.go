package config

type DB struct {
	Name           string `yaml:"name,omitempty"`
	Host           string `yaml:"host,omitempty"`
	User           string `yaml:"user,omitempty"`
	Password       string `yaml:"password,omitempty"`
	Port           string `yaml:"port,omitempty"`
	Type           string `yaml:"type,omitempty"`
	RequestTimeout string `yaml:"requestTimeout,omitempty"`
}

type Service struct {
	Query       string `yaml:"query,omitempty"`
	ServiceName string `yaml:"service_name,omitempty"`
	DB          DB     `yaml:"db,omitempty"`
	ServiceType string `yaml:"service_type,omitempty"`
	Params      Params `yaml:"parameters,omitempty"`
	Platform    string `yaml:"platform,omitempty"`
}

type Params struct {
	Type   string                 `yaml:"type,omitempty"`
	Schema []map[string]Validator `yaml:"schema,omitempty"`
}

// Validator is object used to validate URL Param
type Validator struct {
	Required     bool        `yaml:"required,omitempty"`     // Is URL Param required?
	DefaultValue interface{} `yaml:"defaultValue,omitempty"` //  URL Param DefaultValue,mandatory if Required is true
	Type         string      `yaml:"type,omitempty"`         //  URL Param DefaultValue,mandatory if Required is true
}
