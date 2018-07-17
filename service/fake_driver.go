// -*- Mode: Go; indent-tabs-mode: t -*-

package service

import (
	"github.com/edgexfoundry/edgex-go/core/domain/models"
	logger "github.com/edgexfoundry/edgex-go/support/logging-client"
	"github.com/tonyespy/gxds"
)

type FakeDriver struct {
	lc logger.LoggingClient
}

func (s *FakeDriver) DisconnectDevice(address *models.Addressable) error {
	return nil
}

func (s *FakeDriver) Initialize(lc logger.LoggingClient, asyncCh <-chan *gxds.CommandResult) error {
	s.lc = lc
	return nil
}

func (s *FakeDriver) HandleOperation(ro *models.ResourceOperation,
	d *models.Device, do *models.DeviceObject, desc *models.ValueDescriptor,
	value string, send chan<- *gxds.CommandResult) {

	cr := &gxds.CommandResult{RO: ro, Type: gxds.Bool, BoolResult: true}

	send <- cr
	close(send)
}

func (s *FakeDriver) Stop(force bool) error {
	return nil
}
