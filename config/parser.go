package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ErrBadServiceName  = errors.New("bad function name")
	ErrBadQuery        = errors.New("bad sql query")
	ErrBadDBName       = errors.New("bad db name")
	ErrBadDBUser       = errors.New("bad db username")
	ErrBadDBPort       = errors.New("bad db port")
	ErrBadDBHost       = errors.New("bad db host")
	ErrBadDBPassword   = errors.New("bad db password")
	ErrBadURL          = errors.New("bad url")
	ErrBadDefaultValue = errors.New("please provide default value for parameter")
	ErrBadParamsType   = errors.New("bad params type, the correct values are path and body")
	ErrBadDBType       = errors.New("bad db type, the correct values are mssql ,mysql and postgres")
	ErrBadServiceType  = errors.New("bad service type, the correct values are read and write")
	ErrBadAuthUsername = errors.New("empty username in credentials")
	ErrBadAuthPassword = errors.New("password must be at least 10 characters")
	ErrBadFileName     = errors.New("please provide valid file name")
)

// LoadService loads services configuration from service.yml
func LoadService(path string) (*Service, error) {
	if len(path) < 1 {
		return nil, ErrBadFileName
	}

	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	// parse yml config
	conf := &Service{}
	err = yaml.Unmarshal([]byte(buf), conf)
	if err != nil {
		return nil, err
	}

	// get env var
	// Expand specific env var to avoid removing $1 in the query
	conf.DB.Name = os.Expand(conf.DB.Name, os.Getenv)
	conf.DB.Host = os.Expand(conf.DB.Host, os.Getenv)
	conf.DB.User = os.Expand(conf.DB.User, os.Getenv)
	conf.DB.Password = os.Expand(conf.DB.Password, os.Getenv)
	conf.DB.Type = os.Expand(conf.DB.Type, os.Getenv)
	conf.ServiceName = os.Expand(conf.ServiceName, os.Getenv)

	if err = conf.validate(); err != nil {
		return nil, err
	}
	return conf, err
}

// CreateService will write Service to yml file
func CreateService(s *Service) error {
	data, err := yaml.Marshal(&s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.ServiceName+"-config.yml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// validate octo service
func (s *Service) validate() error {
	switch {
	case len(s.ServiceName) == 0:
		return ErrBadServiceName
	case len(s.Query) == 0:
		return ErrBadQuery
	case len(s.DB.Name) == 0:
		return ErrBadDBName
	case len(s.DB.Password) == 0:
		return ErrBadDBPassword
	case len(s.DB.Host) == 0:
		return ErrBadDBHost
	case len(s.DB.User) == 0:
		return ErrBadDBUser
	case len(s.DB.Port) == 0:
		return ErrBadDBPort
	case s.DB.Type != "postgres" && s.DB.Type != "mssql" && s.DB.Type != "mysql":
		return ErrBadDBType
	case s.ServiceType != "read" && s.ServiceType != "write":
		return ErrBadServiceType
	case len(s.Params.Type) > 0:
		if s.Params.Type != "body" && s.Params.Type != "path" {
			return ErrBadParamsType
		}
	case len(s.Params.Schema) > 0:
		for _, v := range s.Params.Schema { // validate schema
			for _, v := range v {
				if !v.Required && v.DefaultValue == nil {
					return ErrBadDefaultValue
				}
			}
		}
	}
	return nil
}
