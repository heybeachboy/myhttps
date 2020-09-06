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

func BenchmarkCheckData(b *testing.B) {
	 in := []string{"wq","wqewqeqeqeqeqeqe","saeqwygu234639449060570sfdsdfsferwrwrwrwrwrwrwwqewqewqweqweqrwerwrwrwrbnvnzxczdsadaewrwretterwqwqwqw",
	 	"wqweqye211371seweqe34905943570747057375375374573dadwqd457034537sdrksfheuwfe"}
	 out1 := []bool{false,false,false,false}
	 out2 := []bool{true,true,true,true}
	 memory.MemStorage.Flush()
	 for i:=0; i < b.N; i++ {
	 	got := CheckData(in)
	 	if !reflect.DeepEqual(got,out1) && i==0 {
            b.Errorf("got:%v and expect:%v",got,out1)
		}
		 if !reflect.DeepEqual(got,out2) && i>0 {
            b.Errorf("got:%v and expect:%v",got,out2)
		 }
	 }
	 memory.MemStorage.Flush()
}
