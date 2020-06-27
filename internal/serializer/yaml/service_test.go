package yaml

import (
	"fmt"
	"hls-devkit/hlsadmin/internal/configapp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalSetting(t *testing.T) {
	scenarios := []struct {
		input       *configapp.Setting
		expected    []byte
		expectedErr error
	}{
		{
			input: &configapp.Setting{
				Cred: configapp.CredRepo{
					Name: "help",
				},
			},
			expected: []byte(`CredRepo:
  Name: help
`),
			expectedErr: nil,
		},
	}

	for index, scenario := range scenarios {
		svc := &serializeSvc{}
		output, err := svc.MarshalSettings(scenario.input)
		assert.Equal(t, scenario.expectedErr, err, fmt.Sprintf("Scenario: %d. Compare Error", index))
		assert.Equal(t, scenario.expected, output, fmt.Sprintf("Scenario: %d. Compare Expectation", index))
	}
}
