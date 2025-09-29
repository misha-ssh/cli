package cp

type TypeCopy string

const (
	Download TypeCopy = "download"
	Upload   TypeCopy = "upload"
)

// Fields A struct that copy files
type Fields struct {
	TypeCopy   TypeCopy
	LocalFile  string
	RemoteFile string
}
