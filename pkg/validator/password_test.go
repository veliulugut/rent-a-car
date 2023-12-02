package validator

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsPasswordValid_Validate(t *testing.T) {
	tests := []struct {
		password    string
		expectedErr error
	}{
		//Pass
		{
			password:    "123@-abc#",
			expectedErr: nil,
		},

		// did not pass
		{
			password:    "sad",
			expectedErr: ErrrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ValidPassword_%s", tt.password), func(t *testing.T) {
			vld := IsPasswordValid{Password: tt.password}

			err := vld.Validate()

			if !errors.Is(err, tt.expectedErr) {
				t.Error(err)
			}
		})
	}
}
