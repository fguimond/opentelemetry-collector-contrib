// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal/metadata"
)

func TestHealthCheck(t *testing.T) {
	defaultConfig := createDefaultConfig().(*Config)
	defaultConfig.WebHook.Endpoint = "localhost:0"
	consumer := consumertest.NewNop()
	receiver, err := newTracesReceiver(receivertest.NewNopSettings(metadata.Type), defaultConfig, consumer)
	require.NoError(t, err, "failed to create receiver")

	r := receiver
	require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()), "failed to start receiver")
	defer func() {
		require.NoError(t, r.Shutdown(context.Background()), "failed to shutdown receiver")
	}()

	w := httptest.NewRecorder()
	r.handleHealthCheck(w, httptest.NewRequest(http.MethodGet, "http://localhost/health", nil))

	response := w.Result()
	require.Equal(t, http.StatusOK, response.StatusCode)
}
