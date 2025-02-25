/*
Copyright 2020 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package batchscheduler

import (
	"fmt"
	"github.com/spotify/flink-on-k8s-operator/internal/batchscheduler/external"
	"sync"

	"k8s.io/klog"

	schedulerinterface "github.com/spotify/flink-on-k8s-operator/internal/batchscheduler/types"
	"github.com/spotify/flink-on-k8s-operator/internal/batchscheduler/volcano"
)

var (
	mutex            sync.Mutex
	once             = sync.Once{}
	schedulerPlugins = map[string]schedulerinterface.BatchScheduler{}
)

func init() {
	scheduler, err := volcano.New()
	if err != nil {
		klog.Errorf("Failed initializing volcano batch scheduler: %v", err)
		return
	}
	schedulerPlugins[scheduler.Name()] = scheduler

	externalScheduler, err := external.New()
	if err != nil {
		klog.Errorf("Failed initializing external batch scheduler: %v", err)
		return
	}
	schedulerPlugins[externalScheduler.Name()] = externalScheduler
}

// GetScheduler gets the real batch scheduler.
func GetScheduler(name string) (schedulerinterface.BatchScheduler, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if scheduler, exist := schedulerPlugins[name]; exist {
		return scheduler, nil
	}
	return nil, fmt.Errorf("failed to find batch scheduler named with %s", name)
}

func GetRegisteredNames() []string {
	mutex.Lock()
	defer mutex.Unlock()
	var pluginNames []string
	for key := range schedulerPlugins {
		pluginNames = append(pluginNames, key)
	}
	return pluginNames
}
