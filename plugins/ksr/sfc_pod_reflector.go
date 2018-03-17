// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ksr

import (
	"sync"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/golang/protobuf/proto"

	"github.com/contiv/vpp/plugins/ksr/model/pod"
	"github.com/contiv/vpp/plugins/ksr/model/sfc"
)

// PodReflector subscribes to K8s cluster to watch for changes in the
// configuration of k8s pods. Protobuf-modelled changes are published
// into the selected key-value store.
type SfcPodReflector struct {
	Reflector
}

// Init subscribes to K8s cluster to watch for changes in the configuration
// of k8s pods. The subscription does not become active until Start()
// is called.
func (spr *SfcPodReflector) Init(stopCh2 <-chan struct{}, wg *sync.WaitGroup) error {

	podReflectorFuncs := ReflectorFunctions{
		EventHdlrFunc: cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				spr.addPod(obj)
			},
			DeleteFunc: func(obj interface{}) {
				spr.deletePod(obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				spr.updatePod(oldObj, newObj)
			},
		},
		ProtoAllocFunc: func() proto.Message {
			return &sfc.Sfc{}
		},
	}

	return spr.ksrInit(stopCh2, wg, pod.KeyPrefix(), "pods", &coreV1.Pod{}, podReflectorFuncs)
}

// addPod adds state data of a newly created K8s pod into the data store.
func (spr *SfcPodReflector) addPod(obj interface{}) {
	spr.Log.WithField("sfc-pod", obj).Info("sfc-addPod")
	k8sPod, ok := obj.(*coreV1.Pod)
	if !ok {
		spr.Log.Warn("Failed to cast newly created pod object")
		spr.stats.ArgErrors++
		return
	}
	labels := k8sPod.GetLabels()
	if labels != nil {
		for key, val := range labels {
			if key == "sfc" && val == "true" {
				sfcKey := sfc.Key(k8sPod.Name)
				sfcValue := valueToProto(k8sPod.Name, k8sPod.Spec.NodeName)
				spr.ksrAdd(sfcKey, sfcValue)
			}
		}
	}
}

// deletePod deletes state data of a removed K8s pod from the data store.
func (spr *SfcPodReflector) deletePod(obj interface{}) {
	spr.Log.WithField("sfc-pod", obj).Info("sfc-deletePod")
	k8sPod, ok := obj.(*coreV1.Pod)
	if !ok {
		spr.Log.Warn("Failed to cast newly created pod object")
		spr.stats.ArgErrors++
		return
	}
	labels := k8sPod.GetLabels()
	if labels != nil {
		for key, val := range labels {
			if key == "sfc" && val == "true" {
				sfcKey := sfc.Key(k8sPod.Name)
				spr.ksrDelete(sfcKey)
			}
		}
	}
}

// updatePod updates state data of a changed K8s pod in the data store.
func (spr *SfcPodReflector) updatePod(oldObj, newObj interface{}) {
	oldK8sPod, ok1 := oldObj.(*coreV1.Pod)
	newK8sPod, ok2 := newObj.(*coreV1.Pod)
	if !ok1 || !ok2 {
		spr.Log.Warn("Failed to cast changed pod object")
		spr.stats.ArgErrors++
		return
	}

	oldLabels := oldK8sPod.GetLabels()
	newLabels := oldK8sPod.GetLabels()
	oldSfcLabelExist := false
	newSfcLabelExist := false

	if oldLabels != nil {
		for key, val := range oldLabels {
			if key == "sfc" && val == "true" {
				oldSfcLabelExist = true
			}
		}
	}
	if oldLabels != nil {
		for key, val := range newLabels {
			if key == "sfc" && val == "true" {
				newSfcLabelExist = true
			}
		}
	}

	if oldSfcLabelExist && newSfcLabelExist {
		sfcKey := sfc.Key(oldK8sPod.GetName())
		oldSfcValue := valueToProto(oldK8sPod.GetName(), oldK8sPod.Spec.NodeName)
		newSfcValue := valueToProto(newK8sPod.GetName(), newK8sPod.Spec.NodeName)
		spr.ksrUpdate(sfcKey, oldSfcValue, newSfcValue)
	} else if !oldSfcLabelExist && newSfcLabelExist {
		sfcKey := sfc.Key(newK8sPod.GetName())
		sfcValue := valueToProto(newK8sPod.GetName(), newK8sPod.Spec.NodeName)
		spr.ksrAdd(sfcKey, sfcValue)
	} else if oldSfcLabelExist && !newSfcLabelExist {
		sfcKey := sfc.Key(oldK8sPod.GetName())
		spr.ksrDelete(sfcKey)
	}
}

// valueToProto returns the value of the sfc tree a given K8s pod is stored in the
// data store.
func valueToProto(name string, nodeName string) *sfc.Sfc {
	valueProto := &sfc.Sfc{}
	valueProto.Vnf = name
	valueProto.Node = nodeName

	return valueProto
}
