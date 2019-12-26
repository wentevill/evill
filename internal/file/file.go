package file

type File interface {
	Load(interface{}, string) error
}
