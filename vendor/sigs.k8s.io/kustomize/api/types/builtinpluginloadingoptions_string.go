// Code generated by "stringer -type=BuiltinPluginLoadingOptions"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BploUndefined-0]
	_ = x[BploUseStaticallyLinked-1]
	_ = x[BploLoadFromFileSys-2]
}

const _BuiltinPluginLoadingOptions_name = "BploUndefinedBploUseStaticallyLinkedBploLoadFromFileSys"

var _BuiltinPluginLoadingOptions_index = [...]uint8{0, 13, 36, 55}

func (i BuiltinPluginLoadingOptions) String() string {
	if i < 0 || i >= BuiltinPluginLoadingOptions(len(_BuiltinPluginLoadingOptions_index)-1) {
		return "BuiltinPluginLoadingOptions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinPluginLoadingOptions_name[_BuiltinPluginLoadingOptions_index[i]:_BuiltinPluginLoadingOptions_index[i+1]]
}
