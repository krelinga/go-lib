package video

type FileInfo struct {
	Path         string
	HasNfo       bool
	HasTcProfile bool
}

func NfoPath(path string) string {
	return path + ".nfo"
}

func TcProfilePath(path string) string {
	return NfoPath(path) + ".tcprofile"
}
