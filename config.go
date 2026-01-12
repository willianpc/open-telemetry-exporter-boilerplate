// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package boilerplateexporter // import "go.opentelemetry.io/collector/exporter/debugexporter"

import (
	"go.opentelemetry.io/collector/component"
	// "go.opentelemetry.io/collector/config/configtelemetry"
)

// supportedLevels in this exporter's configuration.
// configtelemetry.LevelNone and other future values are not supported.
// var supportedLevels map[configtelemetry.Level]struct{} = map[configtelemetry.Level]struct{}{
// 	configtelemetry.LevelBasic:    {},
// 	configtelemetry.LevelNormal:   {},
// 	configtelemetry.LevelDetailed: {},
// }

// Config defines configuration for debug exporter.
type Config struct {
	Enabled bool `mapstructure:"enabled,omitempty"`

	// prevent unkeyed literal initialization
	_ struct{}
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
	// if _, ok := supportedLevels[cfg.Verbosity]; !ok {
	// 	return fmt.Errorf("verbosity level %q is not supported", cfg.Verbosity)
	// }

	return nil
}
