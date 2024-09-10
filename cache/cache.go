package cache

type Set struct {
	cache map[string]interface{}
}

func New() *Set {
	return &Set{
		cache: make(map[string]interface{}),
	}
}

func (s *Set) Set(key string, value interface{}) {
	s.cache[key] = value
}

func (s *Set) Get(key string) interface{} {
	return s.cache[key]
}

func (s *Set) Delete(key string) {
	delete(s.cache, key)
}
