package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetSingleton02() *Singleton {
	if lazySingleton != nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
