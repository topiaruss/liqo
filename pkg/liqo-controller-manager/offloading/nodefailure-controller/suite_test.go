// Copyright 2019-2025 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodefailurectrl

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/kubectl/pkg/scheme"

	"github.com/liqotech/liqo/pkg/utils/testutil"
)

func TestNodeFailureController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Failure Controller Suite")
}

var _ = BeforeSuite(func() {
	testutil.LogsToGinkgoWriter()
	Expect(corev1.AddToScheme(scheme.Scheme)).To(Succeed())
})
