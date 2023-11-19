package main

import (
	"testing"
)

func TestUnaryNodes(t *testing.T) {
	headNode := UnaryNode{
		"node val",
		nil,
	}

	nextNode := UnaryNode{
		"next val",
		nil,
	}

	headNode.Next = &nextNode

	retStr, err := NodeToStr[UnaryNode](&headNode)

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("\n retVal: %v", retStr)
}

func TestBinaryNodes(t *testing.T) {
	headNode := BinaryNode{
		"node val",
		nil,
		nil,
	}

	nextNode := BinaryNode{
		"next val",
		nil,
		&headNode,
	}

	headNode.Next = &nextNode

	retStr, err := NodeToStr[BinaryNode](&headNode)

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("\n retVal: %v", retStr)
}

func TestMultiNodes(t *testing.T) {
	headNode := MultiNode{
		"node val",
		[]*MultiNode{},
	}

	adjOneNode := MultiNode{
		"adj1 val",
		[]*MultiNode{},
	}

	adjTwoNode := MultiNode{
		"adj2 val",
		[]*MultiNode{},
	}

	headNode.Adj = append(headNode.Adj, []*MultiNode{&adjOneNode, &adjTwoNode}...)

	retStr, err := NodeToStr[MultiNode](&headNode)

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("\n retVal: %v", retStr)
}
