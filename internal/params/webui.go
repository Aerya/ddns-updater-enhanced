package params

import "github.com/qdm12/ddns-updater/internal/provider"

// ParseProviders validates in-memory JSON configuration submitted by the WebUI.
func ParseProviders(configBytes []byte) (
	providers []provider.Provider, warnings []string, err error,
) {
	return extractAllSettings(configBytes)
}
