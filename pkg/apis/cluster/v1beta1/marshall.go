/*
Copyright 2018 Ridecell, Inc.
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

package v1beta1

import (
	"encoding/json"
	"errors"
)

// Intercept JSON decoding and try to deal with "simple" values before giving
// up and assuming it's a full struct. This allows things like:
//
//    config:
//      foo: bar
//      baz: false
//      section:
//        qux: true
//
//
// in a config section. This is all because the Kubernetes codegen machinery
// can't cope with a map[string]interface{}, since it could be some composite
// type, which would break all kinds of things.
func (v *ConfigValue) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		// Wat?
		return err
	}
	ok := v.unmarshalValue(tmp)
	if ok {
		return nil
	}
	// It was something else, hopefully a JSON object.
	sectionVal, ok := tmp.(map[string]interface{})
	if ok {
		for k, tmp := range sectionVal {
			v.Section[k] = &ConfigValue{}
			v.Section[k].unmarshalValue(tmp)
		}

	}
	return errors.New("error decoding JSON")
}

func (v *ConfigValue) unmarshalValue(tmp interface{}) bool {

	boolVal, ok := tmp.(bool)
	if ok {
		v.Bool = &boolVal
		return true
	}
	intVal, ok := tmp.(int)
	if ok {
		v.Int = &intVal
		return true
	}
	floatVal, ok := tmp.(float64)
	if ok {
		v.Float = &floatVal
		return true
	}
	stringVal, ok := tmp.(string)
	if ok {
		v.String = &stringVal
		return true
	}
	return false
}
