/*
Copyright 2022 The Rook Authors. All rights reserved.

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

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCephBlockPoolRadosNamespaceValidateUpdate(t *testing.T) {
	rn := &CephBlockPoolRadosNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rados-name",
		},
		Spec: CephBlockPoolRadosNamespaceSpec{
			BlockPoolName: "pool-name",
		},
	}
	ur := rn.DeepCopy()
	// validate with same pool name
	_, err := ur.ValidateUpdate(rn)
	assert.NoError(t, err)
	// validate with different pool name
	ur.Spec.BlockPoolName = "new-pool-name"
	_, err = ur.ValidateUpdate(rn)
	assert.Error(t, err)
}