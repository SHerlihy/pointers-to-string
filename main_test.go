package main

import (
	"io/fs"
	"os"
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

	retStr, err := NodeToStr[UnaryNode](&headNode, "")

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

	retStr, err := NodeToStr[BinaryNode](&headNode, "")

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

	retStr, err := NodeToStr[MultiNode](&headNode, "")

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("\n retVal: %v", retStr)
}

func TestSpecifyDest(t *testing.T){
    dest := "/home/thehuge/projects/pointer-structure-to-string-go/test-data/"
    destFS := os.DirFS(dest)
    prevDirMatches, err := fs.Glob(destFS, "str-dir*")
    
	headNode := UnaryNode{
		"node val",
		nil,
	}

	nextNode := UnaryNode{
		"next val",
		nil,
	}

	headNode.Next = &nextNode

	_, err = NodeToStr[UnaryNode](&headNode, dest)
	if err != nil {
		t.Errorf("%v", err)
	}

    curDirMatches, err := fs.Glob(destFS, "str-dir*")
	if err != nil {
		t.Errorf("%v", err)
	}

    if len(curDirMatches) < 1 {
        t.Error("new dir not created")
    }

    if curDirMatches[0] == prevDirMatches[0] {
        t.Error("new dir not created")
    }
}
