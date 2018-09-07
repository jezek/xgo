package xgo

import (
	"io"
	"reflect"
	"testing"
)

func TestKeyActionsFromString(t *testing.T) {
	testCases := []struct {
		s   string
		a   []keyAction
		err error
	}{
		{},
		{
			"a",
			[]keyAction{
				keySymStroke{'a'},
			}, nil,
		},
		{
			"a&b\u0061",
			[]keyAction{
				keySymStroke{'a'},
				keySymStroke{'&'},
				keySymStroke{'b'},
				keySymStroke{0x0061},
			}, nil,
		},
		{
			"%",
			nil, errActionRead{io.EOF},
		},
		{
			"%.",
			nil, errInvalidAction{'.'},
		},
		{
			"%%",
			[]keyAction{
				keySymStroke{'%'},
			}, nil,
		},
		{
			"%+a",
			[]keyAction{
				keySymDown{0x0061},
			}, nil,
		},
		{
			"%-a",
			[]keyAction{
				keySymUp{0x0061},
			}, nil,
		},
		{
			"%\"a\"",
			[]keyAction{
				keySymStroke{'a'},
			}, nil,
		},
	}

	for _, tc := range testCases {
		a, err := keyActionsFromString(tc.s)
		if !reflect.DeepEqual(a, tc.a) || !reflect.DeepEqual(err, tc.err) {
			t.Errorf("keyActionsFromString(%s) = (%v, %v), want (%v, %v)", tc.s, a, err, tc.a, tc.err)
		}
	}
}
