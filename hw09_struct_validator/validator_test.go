package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "123456789123456789012345678901234567",
				Name:   "Ivan",
				Age:    25,
				Email:  "ivan@sdf.com",
				Role:   "admin",
				Phones: []string{"12345678911"},
			},
			expectedErr: ValidationErrors(nil),
		},
		{
			in: App{
				Version: "12345",
			},
			expectedErr: ValidationErrors(nil),
		},
		{
			in: Token{
				Header:    make([]byte, 0),
				Payload:   make([]byte, 0),
				Signature: make([]byte, 0),
			},
			expectedErr: ValidationErrors(nil),
		},
		{
			in:          Response{Code: 200},
			expectedErr: ValidationErrors(nil),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestInvalidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "12345678912345678901234567890123456723",
				Name:   "Ivan",
				Age:    17,
				Email:  "ivansdf.com",
				Role:   "admin",
				Phones: []string{"12345678911"},
			},
			expectedErr: fmt.Errorf(
				"field: ID | err: len must be equal to 36, current: 38\n" +
					"field: Age | err: value must be no less then 18, current: 17\n" +
					"field: Email | err: regexp mismatch error - ^\\w+@\\w+\\.\\w+$, current: ivansdf.com\n"),
		},
		{
			in: App{
				Version: "1",
			},
			expectedErr: fmt.Errorf("field: Version | err: len must be equal to 5, current: 1" + "\n"),
		},
		{
			in:          Response{Code: 201},
			expectedErr: fmt.Errorf("field: Code | err: in error number 201 not contained in 200,404,500" + "\n"),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.EqualError(t, tt.expectedErr, err.Error())
		})
	}
}
