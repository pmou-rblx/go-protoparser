package protoparser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseRPC(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		wantRPC           *RPC
		wantRecentScanned string
		wantErr           bool
	}{
		{
			name:    "parse an empty",
			wantErr: true,
		},
		{
			name:  "parse a normal rpc",
			input: "rpc Search (SearchRequest) returns (SearchResponse);",
			wantRPC: &RPC{
				Name: "Search",
				Argument: &Type{
					Name: "SearchRequest",
				},
				Return: &Type{
					Name: "SearchResponse",
				},
			},
		},
		{
			name:  "parse a normal rpc with the emptyStatement option",
			input: "rpc Search (SearchRequest) returns (SearchResponse) {}",
			wantRPC: &RPC{
				Name: "Search",
				Argument: &Type{
					Name: "SearchRequest",
				},
				Return: &Type{
					Name: "SearchResponse",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			lex := newlexer(strings.NewReader(test.input))
			got, err := parseRPC(lex)
			switch {
			case test.wantErr:
				if err == nil {
					t.Errorf("got err nil, but want err")
				}
				return
			case !test.wantErr && err != nil:
				t.Errorf("got err %v, but want nil", err)
				return
			}

			if !reflect.DeepEqual(got, test.wantRPC) {
				t.Errorf("got %v, but want %v", got, test.wantRPC)
			}
			if lex.text() != test.wantRecentScanned {
				t.Errorf("got %v, but want %v", lex.text(), test.wantRecentScanned)
			}
		})
	}
}