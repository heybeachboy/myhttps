package memory
import (
	"sync"
	"sync/atomic"
)

var MemStorage * Storage
func init() {
     MemStorage = NewStorage()
}
func NewStorage()*Storage {
	 return &Storage{total: 0,data: new(sync.Map)}
}
type Storage struct {
	 total uint32
	 data *sync.Map
}
/**
 *add new data
 */
func(s *Storage)Push(key , val interface{}) {
	s.data.Store(key, val)
	atomic.AddUint32(&s.total,1)
}

/**
 *delete data by key
 */
func (s *Storage)Remove(key interface{}) {
   s.data.Delete(key)
   atomic.AddUint32(&s.total,uint32(int32(1)))
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
     s.data.Delete(k)
     return true
}
/**
 *get all data of count
 */
func (s *Storage)Size()uint32 {
	return s.total
}