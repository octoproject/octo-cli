package faas

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ErrBadRequest   = errors.New("bad request")
	ErrBadImageName = errors.New("bad image name provided valid name using --image flag")
)

const (
	FUNCTIONS_ENDPOINT = "/system/functions"
	FUNCTION_ENDPOINT  = "/system/function"
	defaultTimeout     = 40 * time.Second
)

type FaasClient struct {
	Username string
	Password string
	Gateway  string
	client   *http.Client
}

type Function struct {
	ServiceName string            `json:"service"`
	Image       string            `json:"image"`
	Network     string            `json:"network"`
	EnvProcess  string            `json:"envProcess"`
	EnvVars     map[string]string `json:"envVars"`
	Namespace   string            `json:"namespace,omitempty"`
}

func New(username, password, gateway string) *FaasClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: defaultTimeout,
	}
	return &FaasClient{
		Username: username,
		Password: password,
		Gateway:  gateway,
		client: &http.Client{
			Transport: tr,
			Timeout:   time.Second * 5,
		},
	}
}

//DeployFunction will create or update function if it exists
func (c *FaasClient) DeployFunction(f *Function) error {
	serviceExists, err := c.serviceExists(f.ServiceName)
	if err != nil {
		return err
	}

	if !serviceExists {
		err = c.deploy(f, false)
		if err != nil {
			return err
		}
		return nil
	}
	//update function
	err = c.deploy(f, true)
	if err != nil {
		return err
	}
	return nil
}

func (c *FaasClient) serviceExists(name string) (bool, error) {
	code, err := c.getService(name)
	if err != nil && code == http.StatusNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (c *FaasClient) getService(name string) (int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", c.Gateway, FUNCTION_ENDPOINT, name), nil)
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		out, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}
		return res.StatusCode, fmt.Errorf("Status Code: %d, message: %s\n", res.StatusCode, string(out))
	}
	return res.StatusCode, nil
}

func (c *FaasClient) deploy(f *Function, update bool) error {
	if len(f.Image) < 1 {
		return ErrBadImageName
	}

	method := http.MethodPost
	if update {
		method = http.MethodPut
	}

	reqBytes, _ := json.Marshal(&f)
	reader := bytes.NewReader(reqBytes)

	req, err := http.NewRequest(method, c.Gateway+FUNCTIONS_ENDPOINT, reader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		out, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Status Code: %d, message: %s\n", res.StatusCode, string(out))
	}
	return nil
}
