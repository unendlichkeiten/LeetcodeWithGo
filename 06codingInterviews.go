package leetCode

import (
	"log"
)



func travelTail2Head(list LinkList) {
	if nil == list {
		log.Println("node is nil")
		return
	} else {
		travelTail2Head(list.next)
	}
	log.Printf("value %03d", list.data)
}

func travelTail2HeadByReverse(list LinkList) {
	if nil == list {
		log.Println("node is nil")
		return
	}

	// reverse list
	tmp, new, cur := LinkList(nil), LinkList(nil), list
	for ; cur != nil; {
		tmp = cur
		cur = cur.next
		tmp.next = new
		new = tmp
	}

	// travel list
	for ; new != LinkList(nil); {
		log.Printf("value %03d", new.data)
		new = new.next
	}
}