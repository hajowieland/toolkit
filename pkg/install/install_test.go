/*
Copyright 2020 The Flux CD contributors.

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

package install

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	opts := MakeDefaultOptions()
	output, err := Generate(opts)
	if err != nil {
		t.Fatal(err)
	}

	for _, component := range opts.Components {
		img := fmt.Sprintf("%s/%s", opts.Registry, component)
		if !strings.Contains(string(output), img) {
			t.Errorf("component image '%s' not found", img)
		}
	}

	fmt.Println(string(output))
}