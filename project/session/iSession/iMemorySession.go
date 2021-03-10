package iSession

type IMemorySeesion interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (value interface{}, err error)
	Del(key string) (err error)
	Save() (err error)
}
