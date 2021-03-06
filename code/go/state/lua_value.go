package state

import (
	"github.com/hashstone/luago-book/code/go/api"
)

type luaValue interface{}

func typeOf(val luaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64:
		return api.LUA_TNUMBER
	case float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		panic("todo! ")
	}
}

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}
