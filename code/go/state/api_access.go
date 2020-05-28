package state

import (
	"fmt"

	"github.com/hashstone/luago-book/code/go/api"
)

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_typename
func (state *luaState) TypeName(tp api.LuaType) string {
	switch tp {
	case api.LUA_TNONE:
		return "no value"
	case api.LUA_TNIL:
		return "nil"
	case api.LUA_TBOOLEAN:
		return "boolean"
	case api.LUA_TNUMBER:
		return "number"
	case api.LUA_TSTRING:
		return "string"
	case api.LUA_TTABLE:
		return "table"
	case api.LUA_TFUNCTION:
		return "function"
	case api.LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_type
func (state *luaState) Type(idx int) api.LuaType {
	if state.stack.isValid(idx) {
		val := state.stack.get(idx)
		return typeOf(val)
	}
	return api.LUA_TNONE
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isnone
func (state *luaState) IsNone(idx int) bool {
	return state.Type(idx) == api.LUA_TNONE
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isnil
func (state *luaState) IsNil(idx int) bool {
	return state.Type(idx) == api.LUA_TNIL
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isnoneornil
func (state *luaState) IsNoneOrNil(idx int) bool {
	return state.Type(idx) <= api.LUA_TNIL
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isboolean
func (state *luaState) IsBoolean(idx int) bool {
	return state.Type(idx) == api.LUA_TBOOLEAN
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_istable
func (state *luaState) IsTable(idx int) bool {
	return state.Type(idx) == api.LUA_TTABLE
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isfunction
func (state *luaState) IsFunction(idx int) bool {
	return state.Type(idx) == api.LUA_TFUNCTION
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isthread
func (state *luaState) IsThread(idx int) bool {
	return state.Type(idx) == api.LUA_TTHREAD
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isstring
func (state *luaState) IsString(idx int) bool {
	t := state.Type(idx)
	return t == api.LUA_TSTRING || t == api.LUA_TNUMBER
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isnumber
func (state *luaState) IsNumber(idx int) bool {
	_, ok := state.ToNumberX(idx)
	return ok
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_isinteger
func (state *luaState) IsInteger(idx int) bool {
	val := state.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_toboolean
func (state *luaState) ToBoolean(idx int) bool {
	val := state.stack.get(idx)
	return convertToBoolean(val)
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_tointeger
func (state *luaState) ToInteger(idx int) int64 {
	i, _ := state.ToIntegerX(idx)
	return i
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_tointegerx
func (state *luaState) ToIntegerX(idx int) (int64, bool) {
	val := state.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_tonumber
func (state *luaState) ToNumber(idx int) float64 {
	n, _ := state.ToNumberX(idx)
	return n
}

// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_tonumberx
func (state *luaState) ToNumberX(idx int) (float64, bool) {
	val := state.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

// [-0, +0, m]
// http://www.lua.org/manual/5.3/manual.html#lua_tostring
func (state *luaState) ToString(idx int) string {
	s, _ := state.ToStringX(idx)
	return s
}

func (state *luaState) ToStringX(idx int) (string, bool) {
	val := state.stack.get(idx)

	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x) // todo
		state.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
