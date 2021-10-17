package er

type Code uint32

const (
	COpenFile Code = 1 + iota
)

const CodeGap Code = 1 << 16
