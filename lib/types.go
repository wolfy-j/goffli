package lib

import (
	"encoding/json"
	"github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
)

// Encode marshals given interface into value withing lua state (using json bridge)
func Encode(v interface{}, l *lua.LState) (lua.LValue, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	lv, err := luajson.Decode(l, data)
	if err != nil {
		return nil, err
	}

	return lv, nil
}

// Decode unmarshal lua value into given variable using json bridge (make sure to pass pointer!)
func Decode(lv lua.LValue, v interface{}) error {
	data, err := luajson.Encode(lv)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
