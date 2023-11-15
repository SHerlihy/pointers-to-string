package ptrToStr

import (
    "fmt"
)

type UnaryNode interface {
    string
    //*UnaryNode
}

type BinaryNode interface {
    string
    //*BinaryNode
    //*BinaryNode
}

type MultiNode interface {
    string
    //[]*MultiNode
}

type HeadNode interface {
    UnaryNode | BinaryNode | MultiNode
}

func NodeToStr[T HeadNode](head *T) string {
    retStr := ""

    headItem := (*head)

    fmt.Printf("\n headItem %v", headItem)

    //tItem := any(headItem).(T)
    //unaryItem := UnaryNode(tItem)

    //fmt.Printf("\n unaryItem %v", unaryItem)

//    unaryHeadItem := UnaryNode(*head)
//    fmt.Printf("\n unaryHeadItem %v", unaryHeadItem)
//
//    retStr = unaryNodesToString(&unaryHeadItem)
//    switch j:=head.(type) {
//        // fmt.Printf("\n j %v", j)
//    case nil:
//        // err
//    case *UnaryNode:
//        //func unary string
//        retStr = unaryNodesToString(head.(*UnaryNode))
//    case BinaryNode:
//        //func binary string
//    case MultiNode:
//        //func multi string
//    default:
//        //err
//        fmt.Printf("\n deef case %v", head)
//        fmt.Printf("\n deef type %v", j)
//    }

    return retStr
}

func unaryNodesToString[T UnaryNode](head *T) string {
    graphString:="graph TD\n"

    fmt.Printf("\n headItem %v", (*head))

//    func(){
//        from := fmt.Sprintf("%s[%s] --> ", head.Val, head.Val)
//        head = head.Next
//
//        if head == nil {
//            return
//        }
//
//        to := fmt.Sprintf("%s[%s]\n", head.Val, head.Val)
//
//        edge := fmt.Sprintf("%s%s", from, to)
//        graphString = fmt.Sprintf("%s%s", graphString, edge)
//    }()

    return graphString
}
