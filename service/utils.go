package service

import (
	"fmt"

	consulclient "github.com/edgexfoundry/edgex-go/support/consul-client"
	"github.com/tonyespy/gxds"
)

func ConnectToConsul(DeviceServiceName string, conf *gxds.Config) error {

	// Initialize service on Consul
	err := consulclient.ConsulInit(consulclient.ConsulConfig{
		ServiceName:    DeviceServiceName,
		ServicePort:    conf.Service.Port,
		ServiceAddress: conf.Service.Host,
		CheckAddress:   conf.Registry.CheckAddress,
		CheckInterval:  conf.Registry.CheckInterval,
		ConsulAddress:  conf.Registry.Host,
		ConsulPort:     conf.Registry.Port,
	})

	if err != nil {
		return fmt.Errorf("connection to Consul could not be made: %v", err.Error())
	} /* else {
		// Update configuration data from Consul
		/*if err := consulclient.CheckKeyValuePairs(&conf, internal.CoreDataServiceKey, strings.Split(conf.ConsulProfilesActive, ";")); err != nil {
			return fmt.Errorf("error getting key/values from Consul: %v", err.Error())
		}
	}*/
	return nil
}
