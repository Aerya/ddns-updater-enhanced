package provider_test

import (
	"testing"

	"github.com/qdm12/ddns-updater/internal/provider"
	"github.com/qdm12/ddns-updater/internal/provider/constants"
)

func TestAllProviderChoicesHaveWebUIDefinitions(t *testing.T) {
	t.Parallel()

	for _, providerID := range constants.ProviderChoices() {
		if _, ok := provider.Definitions[string(providerID)]; !ok {
			t.Errorf("provider %q has no WebUI field definition", providerID)
		}
	}
}
