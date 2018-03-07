// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package cache

import (
	"sync"
	
	"github.com/edgexfoundry/core-domain-go/models"
)

var (
	ocOnce      sync.Once
	objectCache *ObjectCache
)

type ObjectCache struct {
	objects       map[string]map[string][]string
	responses     map[string]map[string][]models.Reading
	cacheSize     int
	transformData bool
}

func NewObjectCache() *ObjectCache {

	ocOnce.Do(func() {
		objectCache = &ObjectCache{}
	})

	return objectCache
}

//   public String get(String deviceId, String object) JsonObject (java) {
func (oc *ObjectCache) Get(device models.Device, op models.ResourceOperation) string {
	return ""
}

func (oc *ObjectCache) Put(device models.Device, op models.ResourceOperation, value string) {
}

func (oc *ObjectCache) GetResponses(device models.Device, op models.ResourceOperation) []models.Reading {
	return nil
}

func (oc *ObjectCache) GetTransformData() bool {
	return false
}

func (oc *ObjectCache) SetTransformData(transform bool) {
}