package bufb

import (
    "errors"
)

type BufEle struct {
    val []byte   
    next *BufEle 
}

type Buf struct {
    head *BufEle // oldest element
    tail *BufEle // most recently inserted element
}

func NewBuf() *Buf {
    return new(Buf)
}

func (bp *Buf) Insert(val []byte) {
    ele := &BufEle{val : val}
    if bp.head == nil {
        // insert into empyt list
        bp.head = ele
        bp.tail = ele
    } else {
        bp.tail.next = ele
        bp.tail = ele 
    }
}

func (bp *Buf) Front() ([]byte, error) {
    if bp.head == nil {
        return nil, errors.New("Empyt Buffer")
    }
    return bp.head.val, nil
}

func (bp *Buf) Remove() ([]byte, error) {
    e := bp.head
    if e == nil {
        return nil, errors.New("Empty Buffer")
    }
    bp.head = e.next
    if e == bp.tail {
        bp.tail = nil
    }
    return e.val, nil
}

func (bp *Buf) Empty() bool {
    return bp.head == nil
}

func (bp *Buf) Flush() {
    bp.head = nil
    bp.tail = nil
}
