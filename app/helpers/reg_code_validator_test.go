package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidRegCode(t *testing.T) {
    type args struct {
        a string
    }

    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "Should return false",
            args: args{a: "12342134"},
            want: false,
        },
        {
            name: "Should return true",
            args: args{a: "39508134221"},
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := IsValidRegCode(tt.args.a)
            assert.Equal(t, tt.want, got)
        })
    }

}
