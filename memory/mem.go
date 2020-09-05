package memory
import "sync"

type Storage struct {
	 data *sync.Map
}
/**
 *add new data
 */
func(s *Storage)Push(key , val interface{}) {
	s.data.Store(key, val)
}

/**
 *delete data by key
 */
func (s *Storage)Remove(key interface{}) {
   s.data.Delete(key)
}

/**
 *@delete all data
 */
func (s *Storage)Flush() {
	s.data.Range(s.rangeCallback)
}
/**
 *the key is exist
 */
func (s *Storage)IsExist(key interface{})bool {
	_,ok :=s.data.Load(key)
	return ok
}

/**
 *range syn map
 */
func (s *Storage)rangeCallback(k,v interface{})bool {
     s.Remove(k)
     return true
}