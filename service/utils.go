package service

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/BurntSushi/toml"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/tonyespy/gxds"
)

var (
	consul *consulapi.Client = nil
)

// Return Consul client instance
func getConsulClient(config *gxds.Config) (*consulapi.Client, error) {
	consulUrl := config.Registry.Host + ":" + strconv.Itoa(config.Registry.Port)
	defaultConfig := consulapi.DefaultConfig()
	defaultConfig.Address = consulUrl

	consul, err := consulapi.NewClient(defaultConfig)
	if err != nil {
		return nil, err
	} else {
		return consul, nil
	}
}

// Register service in Consul and add health check
func registerDeviceService(consul *consulapi.Client, deviceServiceName string, config *gxds.Config) error {
	var err error

	// Register device service
	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		Name:    deviceServiceName,
		Address: config.Service.Host,
		Port:    config.Service.Port,
	})
	if err != nil {
		return err
	}

	// Register the Health Check
	err = consul.Agent().CheckRegister(&consulapi.AgentCheckRegistration{
		Name:      "Health Check: " + deviceServiceName,
		Notes:     "Check the health of the API",
		ServiceID: deviceServiceName,
		AgentServiceCheck: consulapi.AgentServiceCheck{
			HTTP:     config.Registry.CheckAddress,
			Interval: config.Registry.CheckInterval,
		},
	})

	return err
}

// TODO(apopovych) Store config data into Consul
func storeConfigInConsul(consul *consulapi.Client, config *gxds.Config) error {
	return nil
}

func connectToConsul(deviceServiceName string, config *gxds.Config) error {
	var err error

	consul, err = getConsulClient(config)
	if err != nil {
		return err
	}

	err = registerDeviceService(consul, deviceServiceName, config)
	if err != nil {
		return err
	}
	err = storeConfigInConsul(consul, config)

	return err
}

// LoadConfig loads the local configuration file based upon the
// specified parameters and returns a pointer to the global Config
// struct which holds all of the local configuration settings for
// the DS.
func loadConfig(profile string, configDir string) (config *gxds.Config, err error) {
	var name string

	if len(configDir) == 0 {
		configDir = "./res/"
	}

	if len(profile) > 0 {
		name = "configuration-" + profile + ".toml"
	} else {
		name = "configuration.toml"
	}

	path := configDir + name

	// As the toml package can panic if TOML is invalid,
	// or elements are found that don't match members of
	// the given struct, use a defered func to recover
	// from the panic and output a useful error.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("could not load configuration file; invalid TOML (%s)", path)
		}
	}()

	config = &gxds.Config{}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not load configuration file (%s): %v", path, err.Error())
	}

	// Decode the configuration from TOML
	//
	// TODO: invalid input can cause a SIGSEGV fatal error (INVESTIGATE)!!!
	//       - test missing keys, keys with wrong type, ...
	err = toml.Unmarshal(contents, config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration file (%s): %v", path, err.Error())
	}

	return config, nil
}
