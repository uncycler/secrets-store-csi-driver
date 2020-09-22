/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

import (
	"fmt"

	"sigs.k8s.io/secrets-store-csi-driver/apis/v1alpha1"

	"k8s.io/client-go/tools/cache"
)

// SecretProviderClassPodStatusLister is a store to list secretproviderclasspodstatuses
type SecretProviderClassPodStatusLister struct {
	cache.Store
}

// GetWithKey returns secret provider class pod status with key from the informer cache
func (spcpsl *SecretProviderClassPodStatusLister) GetWithKey(key string) (*v1alpha1.SecretProviderClassPodStatus, error) {
	s, exists, err := spcpsl.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("secret provider class pod status not found in informer cache")
	}
	spcps, ok := s.(*v1alpha1.SecretProviderClassPodStatus)
	if !ok {
		return nil, fmt.Errorf("failed to cast %T to %s", s, "secretproviderclasspodstatus")
	}
	return spcps, nil
}
