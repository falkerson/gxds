// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
//
// SPDX-License-Identifier: Apache-2.0
//
package service

import (
	"testing"

	"github.com/edgexfoundry/edgex-go/core/domain/models"
)

// TODO:
//   TestCompareDevices
//   TestCompareDeviceProfiles
//   TestCompareDeviceResources
//   TestCompareResources
//   TestCompareResourceOperations
//   TestCompareServices

func TestCompareCommands(t *testing.T) {
	command_1 := &models.Command{Name: "test1"}
	command_2 := &models.Command{Name: "test2"}
	command_3 := &models.Command{Name: "test3"}
	commands_array_1 := []models.Command{*command_1, *command_2}
	commands_array_2 := []models.Command{*command_1, *command_2}
	commands_array_3 := []models.Command{*command_1, *command_3}

	if !compareCommands(commands_array_1, commands_array_2) {
		t.Error("Arrays of commands should be equal")
	}

	if compareCommands(commands_array_1, commands_array_3) {
		t.Error("Arrays of commands should not be equal")
	}
}

func TestCompareDevices(t *testing.T) {

}
