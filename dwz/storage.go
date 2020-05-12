package main

type Storage interface {
	Shorten(url string, exp int64) (string, error) // 传入长链接 有效时间 返回短链接
	ShortlinkInfo(eid string) (interface{}, error) // 传入短链接  返回短链接详情
	Unshorten(eid string) (string, error)          // 传入短链接 返回长链接
}
