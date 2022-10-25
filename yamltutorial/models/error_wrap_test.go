package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintRandomAPIError(t *testing.T) {
	tests := []struct {
		name string
		want *APIError
	}{
		{
			name: "random",
			want: &APIError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintRandomAPIError(); !reflect.DeepEqual(got, tt.want) {
				// json.Marshal(got)
				// fmt.Println(got.MarshalBinary())
				fmt.Printf("%+v\n", got)
				// t.Errorf("PrintRandomAPIError() = %v, want %v", got, tt.want)
			}
		})
	}

}
