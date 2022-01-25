// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
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

package stack

import (
	_ "embed"
	"testing"
)

func TestFunctionVersionString(t *testing.T) {
	tests := []struct {
		name        string
		funcVersion string
		want        string
	}{
		{
			name: "from embed",
			want: "v0.13.0-rc.11",
		},
		{
			name:        "from function",
			funcVersion: "v0.12.0",
			want:        "v0.12.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				Version: tt.funcVersion,
			}
			if got := f.VersionString(nil); got != tt.want {
				t.Errorf("Function.VersionString() = '%s', want '%s'", got, tt.want)
			}
		})
	}
}
