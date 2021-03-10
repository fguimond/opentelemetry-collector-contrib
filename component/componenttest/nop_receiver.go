// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package componenttest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenthelper"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer"
)

// nopReceiverFactory is factory for nopReceiver.
type nopReceiverFactory struct{}

var nopReceiverFactoryInstance = &nopReceiverFactory{}

// NewNopReceiverFactory returns a component.ReceiverFactory that constructs nop exporters.
func NewNopReceiverFactory() component.ReceiverFactory {
	return nopReceiverFactoryInstance
}

// Type gets the type of the Receiver config created by this factory.
func (f *nopReceiverFactory) Type() configmodels.Type {
	return "nop"
}

// CreateDefaultConfig creates the default configuration for the Receiver.
func (f *nopReceiverFactory) CreateDefaultConfig() configmodels.Receiver {
	return &configmodels.ReceiverSettings{
		TypeVal: f.Type(),
	}
}

// CreateTracesReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateTracesReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ configmodels.Receiver,
	_ consumer.TracesConsumer,
) (component.TracesReceiver, error) {
	return nopReceiverInstance, nil
}

// CreateMetricsReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateMetricsReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ configmodels.Receiver,
	_ consumer.MetricsConsumer,
) (component.MetricsReceiver, error) {
	return nopReceiverInstance, nil
}

// CreateMetricsReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateLogsReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ configmodels.Receiver,
	_ consumer.LogsConsumer,
) (component.LogsReceiver, error) {
	return nopReceiverInstance, nil
}

var nopReceiverInstance = &nopReceiver{
	Component: componenthelper.NewComponent(componenthelper.DefaultComponentSettings()),
}

// nopReceiver stores consumed traces and metrics for testing purposes.
type nopReceiver struct {
	component.Component
}
