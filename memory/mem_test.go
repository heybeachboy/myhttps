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
