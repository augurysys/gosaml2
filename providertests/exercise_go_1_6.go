// +build !go1.7

package providertests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ExerciseProviderTestScenarios(t *testing.T, scenarios []ProviderTestScenario) {
	for _, scenario := range scenarios {
		_, err := scenario.ServiceProvider.RetrieveAssertionInfo(scenario.Response)
		if scenario.CheckError != nil {
			scenario.CheckError(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
