package core

import (
	"myhttps/memory"
	"reflect"
	"testing"
)

func TestCheckData(t *testing.T) {
	 tests:= []struct {
	 	in []string
	 	out []bool

	 }{
	 	 {[]string{"qw","wqwqw","wqw","123"},[]bool{false,false,false,false}},
		 {[]string{"12","22","33"},[]bool{false,false,false}},

	 }
	 memory.MemStorage.Flush()
	 for _,val := range tests {
	 	got := CheckData(val.in)
        if !reflect.DeepEqual(got,val.out) {
        	t.Errorf("got:%v and expect:%v",got,val.out)
		}

	 }
	 memory.MemStorage.Flush()
}
