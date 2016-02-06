package bufb

import (
    "encoding/json"
    "math/rand"
    "testing"
)

func i2b(i int) []byte {
    b, _ := json.Marshal(i)
    return b
}

func b2i(b []byte) int {
    var i int
    json.Unmarshal(b, &i)
    return i
}

var ntest int = 10

var nele int = 50

func runtest(t *testing.T, bp *Buf) {
    inserted := 0
    removed := 0
    emptycount := 0
    for removed < nele {
        if bp.Empty() {
            emptycount++
        }
        insert := !(inserted == nele)
        if inserted > removed && rand.Int31n(2) == 0 {
            insert = false
        }
        if insert {
            bp.Insert(i2b(inserted))
            inserted++
        } else {
            b, err := bp.Remove()
            if err != nil {
                t.Logf("Attempt to remove from empyt buffer \n")
                t.Fail()
            }
            v := b2i(b)
            if v!= removed {
                t.Logf("Removed %d ,Expected %d\n", v, removed)
                t.Fail()
            }
            removed++
        }
    }
}

func TestBuf(t *testing.T) {
    for i := 0; i < ntest; i++ {
        bp := NewBuf()
        runtest(t, bp)
        if !bp.Empty() {
            t.Logf("Expected empyt buffer")
            t.Fail()
        }
    }
}