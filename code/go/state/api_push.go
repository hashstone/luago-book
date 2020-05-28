package state

// [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushnil
func (state *luaState) PushNil() {
	state.stack.push(nil)
}

// [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushboolean
func (state *luaState) PushBoolean(b bool) {
	state.stack.push(b)
}

// [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushinteger
func (state *luaState) PushInteger(n int64) {
	state.stack.push(n)
}

// [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushnumber
func (state *luaState) PushNumber(n float64) {
	state.stack.push(n)
}

// [-0, +1, m]
// http://www.lua.org/manual/5.3/manual.html#lua_pushstring
func (state *luaState) PushString(s string) {
	state.stack.push(s)
}
