package main

import (
	"reflect"
	"testing"
)

func TestCreateCombinationVariants(t *testing.T) {
	tests := []struct {
		emailAddress string
		want         [][2]string
		wantErr      bool
	}{
		{
			emailAddress: "brinks@there.com",
			want: [][2]string{
				{"rinks", ""},
				{"brink", ""},
				{"s", "brink"},
				{"brink", "s"},
				{"rinks", "b"},
				{"b", "rinks"},
			},
		},
		{
			emailAddress: "",
			wantErr:      true,
		},
		{
			emailAddress: "invalid-email",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		got, err := createCombinationVariants(tt.emailAddress)
		if (err != nil) != tt.wantErr {
			t.Errorf("createCombinationVariants() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("createCombinationVariants() = %v, want %v", got, tt.want)
		}
	}
}
