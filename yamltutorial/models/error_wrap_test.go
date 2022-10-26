package models

import (
	"encoding/json"
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
				byteArray, err := got.MarshalBinary()
				if err != nil {
					fmt.Println("Error Occured")
				}
				fmt.Printf("%s", byteArray)
				// t.Errorf("PrintRandomAPIError() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestPrint(t *testing.T) {
	apiErr := PrintRandomAPIError()
	byteArray, err := json.Marshal(apiErr)
	if err != nil {
		fmt.Println("Error Occured")
	}
	fmt.Printf("%s", byteArray)
}
