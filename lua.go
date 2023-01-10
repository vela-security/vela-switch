package vswitch

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

/*
	local switch = vela.switch("name")
	local s = switch._{
		["baidu.com"] = pipe,
	}

	s.once(true)
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.case("name eq baidu.com" , "icmp > 188").pipe()
	s.default(pipe1 , pipe2 , pipe3)
	s.match(app)
*/

func newLuaSwitchL(L *lua.LState) int {
	s := NewSwitchL(L)
	L.Push(s)
	return 1
}

func WithEnv(env assert.Environment) {
	xEnv = env
	xEnv.Set("switch", lua.NewFunction(newLuaSwitchL))
}
