package systray

import "unsafe"

// notifyIconData must match NOTIFYICONDATAW byte for byte: Shell_NotifyIcon
// reads it by offset and checks cbSize against the sizes it knows. These
// indexes are constant and only 0 is in range, so a layout that drifts stops
// the build instead of failing at run time. It did once, silently: with
// uTimeout and uVersion -- a union in C -- as two fields, szInfoTitle and
// everything after it sat 4 bytes late, and Windows took NIM_MODIFY but
// showed no balloon.
var _ = [1]struct{}{}[
// szInfo is 256 WCHARs; the uTimeout/uVersion union follows, then szInfoTitle.
unsafe.Offsetof(notifyIconData{}.InfoTitle)-(unsafe.Offsetof(notifyIconData{}.Info)+256*2+4)]

var _ = [1]struct{}{}[
// hBalloonIcon is the last member.
unsafe.Sizeof(notifyIconData{})-(unsafe.Offsetof(notifyIconData{}.BalloonIcon)+unsafe.Sizeof(uintptr(0)))]

var _ = [1]struct{}{}[
// 976 bytes on 64-bit Windows, 956 on 32-bit.
unsafe.Sizeof(notifyIconData{})-(956+20*(unsafe.Sizeof(uintptr(0))/8))]
