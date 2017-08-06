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
				keySymDown{'a'},
				keySymUp{'a'},
			}, nil,
		},
		{
			"a&b\u0061",
			[]keyAction{
				keySymDown{'a'}, keySymUp{'a'},
				keySymDown{'&'}, keySymUp{'&'},
				keySymDown{'b'}, keySymUp{'b'},
				keySymDown{0x0061}, keySymUp{'a'},
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
				keySymDown{'%'},
				keySymUp{'%'},
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
				keySymDown{0x0061},
			}, nil,
		},
		{
			"%\"a\"",
			[]keyAction{
				keySymDown{'a'}, keySymUp{'a'},
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
