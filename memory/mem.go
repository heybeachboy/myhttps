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
	 list []string
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
	s.data.Range(s.rangeDelCallback)
}
/**
 *the key is exist
 */
func (s *Storage)IsExist(key interface{})bool {
	_,ok :=s.data.Load(key)
	return ok
}

/**
 *get all data
 */
func (s *Storage)GetAllData()[]string {
	 s.list = []string{}
	 s.data.Range(s.rangeGetCallback)
	 return s.list
}

func (s *Storage)rangeGetCallback(k,v interface{})bool {
	  val,ok := v.(string)
	  if ok {
		  s.list = append(s.list,val)
	  }
	  return true
}

/**
 *range syn map
 */
func (s *Storage)rangeDelCallback(k,v interface{})bool {
     s.data.Delete(k)
     return true
}
/**
 *get all data of count
 */
func (s *Storage)Size()uint32 {
	return s.total
}