package idl_test

import (
	"github.com/aviate-labs/candid-go/idl"
)

func ExampleBool() {
	test([]idl.Type{new(idl.BoolType)}, []interface{}{true})
	test([]idl.Type{new(idl.BoolType)}, []interface{}{false})
	test([]idl.Type{new(idl.BoolType)}, []interface{}{0})
	test([]idl.Type{new(idl.BoolType)}, []interface{}{"false"})
	// Output:
	// 4449444c00017e01
	// 4449444c00017e00
	// enc: invalid argument: 0
	// enc: invalid argument: false
}
