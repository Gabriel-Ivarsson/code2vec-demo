// Code generated by "stringer"; DO NOT EDIT.

package sqlproxyccl

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[codeNone-0]
	_ = x[codeAuthFailed-1]
	_ = x[codeBackendReadFailed-2]
	_ = x[codeBackendWriteFailed-3]
	_ = x[codeClientReadFailed-4]
	_ = x[codeClientWriteFailed-5]
	_ = x[codeUnexpectedInsecureStartupMessage-6]
	_ = x[codeUnexpectedStartupMessage-7]
	_ = x[codeParamsRoutingFailed-8]
	_ = x[codeBackendDown-9]
	_ = x[codeBackendRefusedTLS-10]
	_ = x[codeBackendDisconnected-11]
	_ = x[codeClientDisconnected-12]
	_ = x[codeProxyRefusedConnection-13]
	_ = x[codeExpiredClientConnection-14]
	_ = x[codeUnavailable-15]
}

const _errorCode_name = "codeNonecodeAuthFailedcodeBackendReadFailedcodeBackendWriteFailedcodeClientReadFailedcodeClientWriteFailedcodeUnexpectedInsecureStartupMessagecodeUnexpectedStartupMessagecodeParamsRoutingFailedcodeBackendDowncodeBackendRefusedTLScodeBackendDisconnectedcodeClientDisconnectedcodeProxyRefusedConnectioncodeExpiredClientConnectioncodeUnavailable"

var _errorCode_index = [...]uint16{0, 8, 22, 43, 65, 85, 106, 142, 170, 193, 208, 229, 252, 274, 300, 327, 342}

func (i errorCode) String() string {
	if i < 0 || i >= errorCode(len(_errorCode_index)-1) {
		return "errorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _errorCode_name[_errorCode_index[i]:_errorCode_index[i+1]]
}
