// Code generated by "stringer -type=sessionType"; DO NOT EDIT.

package bot

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[sessionUnknown-0]
	_ = x[sessionNormal-1]
	_ = x[sessionRepeat-2]
	_ = x[sessionAutoreply-3]
	_ = x[sessionSubNotification-4]
}

const _sessionType_name = "sessionUnknownsessionNormalsessionRepeatsessionAutoreplysessionSubNotification"

var _sessionType_index = [...]uint8{0, 14, 27, 40, 56, 78}

func (i sessionType) String() string {
	if i < 0 || i >= sessionType(len(_sessionType_index)-1) {
		return "sessionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _sessionType_name[_sessionType_index[i]:_sessionType_index[i+1]]
}
