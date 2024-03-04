/*
The MIT License (MIT)

Copyright (c) 2022-2024 Utility Warehouse

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// XdsServiceSpecRetryApplyConfiguration represents an declarative configuration of the XdsServiceSpecRetry type for use
// with apply.
type XdsServiceSpecRetryApplyConfiguration struct {
	RetryOn      []string                                            `json:"retryOn,omitempty"`
	NumRetries   *uint32                                             `json:"numRetries,omitempty"`
	RetryBackOff *XdsServiceSpecRetryBackoffPolicyApplyConfiguration `json:"backoff,omitempty"`
}

// XdsServiceSpecRetryApplyConfiguration constructs an declarative configuration of the XdsServiceSpecRetry type for use with
// apply.
func XdsServiceSpecRetry() *XdsServiceSpecRetryApplyConfiguration {
	return &XdsServiceSpecRetryApplyConfiguration{}
}

// WithRetryOn adds the given value to the RetryOn field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RetryOn field.
func (b *XdsServiceSpecRetryApplyConfiguration) WithRetryOn(values ...string) *XdsServiceSpecRetryApplyConfiguration {
	for i := range values {
		b.RetryOn = append(b.RetryOn, values[i])
	}
	return b
}

// WithNumRetries sets the NumRetries field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NumRetries field is set to the value of the last call.
func (b *XdsServiceSpecRetryApplyConfiguration) WithNumRetries(value uint32) *XdsServiceSpecRetryApplyConfiguration {
	b.NumRetries = &value
	return b
}

// WithRetryBackOff sets the RetryBackOff field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetryBackOff field is set to the value of the last call.
func (b *XdsServiceSpecRetryApplyConfiguration) WithRetryBackOff(value *XdsServiceSpecRetryBackoffPolicyApplyConfiguration) *XdsServiceSpecRetryApplyConfiguration {
	b.RetryBackOff = value
	return b
}
