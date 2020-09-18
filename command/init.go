package command

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/octoproject/octo-cli/config"
	"github.com/octoproject/octo-cli/prompt"
	"github.com/spf13/cobra"
)

const (
	postgresBindVars = "$"
	mssqlBindVars    = "?"
	mysqlBindVars    = "$"
)

func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate service configuration YAML file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInit()
		},
		SilenceUsage: true,
	}

	return cmd
}

func runInit() error {
	name, err := prompt.PromptString("Enter the service name?", true)
	if err != nil {
		return err
	}

	err = validateServiceName(name)
	if err != nil {
		return err
	}

	ty, err := prompt.PromptSelect("Select the service type:", []string{"read", "write"})
	if err != nil {
		return err
	}

	query, err := prompt.PromptString("Add the SQL query:", true)
	if err != nil {
		return err
	}

	db, err := promptDBQuestions()
	if err != nil {
		return err
	}

	bindvars := getBindVars(db.Type)
	numParam := strings.Count(query, bindvars)

	p, err := promptParams(numParam)
	if err != nil {
		return err
	}

	platform, err := prompt.PromptSelect("Select the platform:", []string{"knative", "openfaas"})
	if err != nil {
		return err
	}

	s := &config.Service{
		ServiceName: name,
		ServiceType: ty,
		Query:       query,
		Params:      *p,
		Platform:    platform,
		DB:          *db,
	}
	return config.CreateService(s)
}

func promptParams(numParam int) (*config.Params, error) {
	if numParam < 0 {
		return nil, nil
	}

	ptype, err := prompt.PromptSelect("Enter the parameters type?", []string{"body", "path"})
	if err != nil {
		return nil, err
	}

	schema := make([]map[string]config.Validator, numParam)
	for i := 0; i < numParam; i++ {
		pname, err := prompt.PromptString("Enter "+strconv.Itoa(i+1)+" parameter:", true)
		if err != nil {
			return nil, err
		}

		isRequired, err := prompt.PromptConfirm("Is this parameter required?")
		if err != nil {
			return nil, err
		}

		var defaultVal string
		if !isRequired {
			defaultVal, err = prompt.PromptString("What is the default value for parameter:", true)
			if err != nil {
				return nil, err
			}
		}

		v := make(map[string]config.Validator, numParam)
		v[pname] = config.Validator{Required: isRequired, DefaultValue: defaultVal}
		schema[i] = v
	}

	return &config.Params{Type: ptype, Schema: schema}, nil
}

func promptDBQuestions() (*config.DB, error) {
	dbDialect, err := prompt.PromptSelect("Select the database dialect:", []string{"postgres", "mssql", "mysql"})
	if err != nil {
		return nil, err
	}

	name, err := prompt.PromptString("Enter the database name:", true)
	if err != nil {
		return nil, err
	}

	host, err := prompt.PromptString("Enter the database host:", true)
	if err != nil {
		return nil, err
	}

	user, err := prompt.PromptString("Enter the database user:", true)
	if err != nil {
		return nil, err
	}

	port, err := prompt.PromptString("Enter the database port:", true)
	if err != nil {
		return nil, err
	}

	pass, err := prompt.PromptString("Enter the database password:", true)
	if err != nil {
		return nil, err
	}

	timeout, err := prompt.PromptString("Enter the database request timeout:", false)
	if err != nil {
		return nil, err
	}

	if len(timeout) < 1 {
		// set default timeout to 30s
		timeout = "30000"
	}

	return &config.DB{
		Name:           name,
		Host:           host,
		User:           user,
		Password:       pass,
		Port:           port,
		Type:           dbDialect,
		RequestTimeout: timeout,
	}, nil
}

func getBindVars(dbDialect string) string {
	switch {
	case dbDialect == "postgres":
		return postgresBindVars
	case dbDialect == "mssql":
		return mssqlBindVars
	case dbDialect == "mysql":
		return mysqlBindVars
	}
	return ""
}

// validateServiceName only allows valid Kubernetes services names
func validateServiceName(serviceName string) error {
	var validDNS = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	if matched := validDNS.MatchString(serviceName); !matched {
		return fmt.Errorf(`Service name can only contain a-z, 0-9 and dashes`)
	}
	return nil
}
