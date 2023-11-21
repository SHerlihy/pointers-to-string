package ptr_to_str

import (
	//	"fmt"
	//	"os"
	"io/fs"
	"fmt"
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

func TestSpecifyDest(t *testing.T) {
    projDir, err := os.Getwd()

	if err != nil {
		t.Errorf("%v", err)
	}

    dest := fmt.Sprintf("%s%s", projDir, "/test-data")
	destFS := os.DirFS(projDir)

	prevDirMatches, err := fs.Glob(destFS, "test-data/str-dir*")

	for i := 0; i < len(prevDirMatches); i++ {
		os.RemoveAll(prevDirMatches[i])
	}

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

	curDirMatches, err := fs.Glob(destFS, "test-data/str-dir*")

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

func TestBinaryTreeToMultNodes(t *testing.T){
	rootNode := TreeNode{
		1,
		nil,
		nil,
	}

	leftNode := TreeNode{
		2,
		nil,
		nil,
	}

	rightNode := TreeNode{
		3,
		nil,
		nil,
	}

	rootNode.Left = &leftNode
	rootNode.Right = &rightNode

	rootMultiRef := BinaryTreeToMultiNodes(&rootNode)

	defStr, err := NodeToStr[MultiNode](rootMultiRef, "")

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("%v", defStr)
}
