package ifile

// 定义接口
type IFile interface {
	// 接口要实现的方法
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

// 读取接口
type IReader interface {
	Read(buf []byte) (n int, err error)
}

// 写入接口
type IWriter interface {
	Write(buf []byte) (n int, err error)
}

// 关闭接口
type ICloser interface {
	Close() error
}

type File struct {
	number int
}

func (f *File) Read(buf []byte) (n int, err error) {
	return n, nil
}
func (f *File) Write(buf []byte) (n int, err error) {
	return n, nil
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return pos, nil
}
func (f *File) Close() error {
	return nil
}
