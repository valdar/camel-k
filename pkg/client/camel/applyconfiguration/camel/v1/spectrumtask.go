/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SpectrumTaskApplyConfiguration represents an declarative configuration of the SpectrumTask type for use
// with apply.
type SpectrumTaskApplyConfiguration struct {
	BaseTaskApplyConfiguration    `json:",inline"`
	PublishTaskApplyConfiguration `json:",inline"`
}

// SpectrumTaskApplyConfiguration constructs an declarative configuration of the SpectrumTask type for use with
// apply.
func SpectrumTask() *SpectrumTaskApplyConfiguration {
	return &SpectrumTaskApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithName(value string) *SpectrumTaskApplyConfiguration {
	b.Name = &value
	return b
}

// WithConfiguration sets the Configuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Configuration field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithConfiguration(value *BuildConfigurationApplyConfiguration) *SpectrumTaskApplyConfiguration {
	b.Configuration = value
	return b
}

// WithContextDir sets the ContextDir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContextDir field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithContextDir(value string) *SpectrumTaskApplyConfiguration {
	b.ContextDir = &value
	return b
}

// WithBaseImage sets the BaseImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseImage field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithBaseImage(value string) *SpectrumTaskApplyConfiguration {
	b.BaseImage = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithImage(value string) *SpectrumTaskApplyConfiguration {
	b.Image = &value
	return b
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *SpectrumTaskApplyConfiguration) WithRegistry(value *RegistrySpecApplyConfiguration) *SpectrumTaskApplyConfiguration {
	b.Registry = value
	return b
}
