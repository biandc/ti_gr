package t_041_232_implement_queue_using_stacks

import (
	"reflect"
	"testing"
)
func TestConstructor(t *testing.T){
    var (
        expected= MyQueue{}
        actual MyQueue
    )
    actual=Constructor()
    if reflect.DeepEqual(expected,actual){
        
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
