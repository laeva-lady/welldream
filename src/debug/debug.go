package debug

var (
	debug = false
)

func Debug() bool {
	return debug
}
func SetDebug(d bool) {
	debug = d
}
