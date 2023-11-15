package ptrToStr

import (
	"fmt"
)

type UnaryNode struct {
	Val string
	Next *UnaryNode
}

type BinaryNode struct {
	Val string
	//*BinaryNode
	//*BinaryNode
}

type MultiNode struct {
	Val string
	//[]*MultiNode
}

type Node interface {
    UnaryNode | BinaryNode | MultiNode
}

func NodeToStr[T Node](head *T) string {
    if head == nil {
        return ""
    }

	retStr := ""

	headItem := (*head)

    switch j:=any(headItem).(type) {
    case UnaryNode:
        retStr = unaryNodesToString(&j)
    case BinaryNode:
        fmt.Printf("Binary")
    case MultiNode:
        fmt.Printf("Multi")
    default:
        fmt.Printf("defautl")
    }

	return retStr
}

func unaryNodesToString(head *UnaryNode) string {
	graphString := "graph TD\n"

	func(){
	    from := fmt.Sprintf("%s[%s] --> ", head.Val, head.Val)
	    head = head.Next
	
	    if head == nil {
	        return
	    }
	
	    to := fmt.Sprintf("%s[%s]\n", head.Val, head.Val)
	
	    edge := fmt.Sprintf("%s%s", from, to)
	    graphString = fmt.Sprintf("%s%s", graphString, edge)
	}()

	return graphString
}
