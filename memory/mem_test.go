package memory

import (
	"testing"
)

func TestStorage_Push(t *testing.T) {
     test := []string{"aa","12","cc","bb","kkk"}
     mem := NewStorage()
     for _,val := range test {
     	mem.Push(val,nil)
	 }
	 for _,val := range test {
	 	if !mem.IsExist(val) {
	 		t.Errorf("%s is not exist",val)
		}
	 }
}

func BenchmarkStorage_Push(b *testing.B) {
	//test := []string{"aawq","21212","ccdfgfdgdg","dgdfgdg","kdgdfgdgkk","qewskusdfjsgd","sdagjkfbdn3298378293497249734"}
	mem := NewStorage()
    for i :=0; i<b.N; i++ {
    	mem.Push("wqeywqioryrhrhwkjq321391842743927349782324908294",nil)
	}
	mem.Flush()
}

func TestStorage_Remove(t *testing.T) {

}

func TestStorage_Flush(t *testing.T) {

}

func TestStorage_IsExist(t *testing.T) {

}
func TestStorage_Size(t *testing.T) {

}

func TestStorage_GetAllData(t *testing.T) {

}
