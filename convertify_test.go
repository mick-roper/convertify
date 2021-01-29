package convertify

import (
	"reflect"
	"testing"
)

type target struct {
	name string
	age  int
}

func TestConvert(t *testing.T) {
	type args struct {
		input  map[string]interface{}
		target interface{}
	}
	tests := []struct {
		name   string
		args   args
		expect target
	}{
		{
			name: "parses a simple type",
			args: args{
				input: map[string]interface{}{
					"name": "something",
					"age":  10,
				},
				target: &target{},
			},
			expect: target{
				name: "something",
				age:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convert(tt.args.input, tt.args.target)

			if !reflect.DeepEqual(tt.args.target, tt.expect) {
				t.Errorf("expect %v but got %v", tt.expect, tt.args.target)
			}
		})
	}
}
