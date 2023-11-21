package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

type UnaryNode struct {
	Val  string
	Next *UnaryNode
}

type BinaryNode struct {
	Val  string
	Next *BinaryNode
	Prev *BinaryNode
}

type MultiNode struct {
	Val string
	Adj []*MultiNode
}

type Node interface {
	UnaryNode | BinaryNode | MultiNode
}

func NodeToStr[T Node](headRef *T, dest string) (string, error) {
	if headRef == nil {
		return "", errors.New("No headRef ref")
	}

	retStr := ""
	var err error

	headItem := (*headRef)

	switch j := any(headItem).(type) {
	case UnaryNode:
		retStr = unaryNodesToString(&j)
	case BinaryNode:
		retStr = binaryNodesToString(&j)
	case MultiNode:
		retStr = multiNodesToString(&j)
	default:
		err = errors.New("head not of type Node")
		return retStr, err
	}

	fmt.Fprintln(os.Stdout, retStr)

    if len(dest) < 1 {
        dest, err = os.Getwd()
        if err != nil {
            return retStr, err
        }
    }

    err = strToTmpFile(retStr, dest)

	return retStr, err
}

func strToTmpFile(str string, dest string)error{
    data := []byte(str)

    projFS := os.DirFS(dest)
    tmpDirMatches, err := fs.Glob(projFS, "str-dir*")

    for i:=0; i<len(tmpDirMatches); i++ {
        os.RemoveAll(tmpDirMatches[i])
    }
    
    tmpDir, err := os.MkdirTemp(dest, "str-dir*")
    temp, err := os.CreateTemp(tmpDir, "pts*.txt")
    if err != nil {
        return err
    }

    if err := temp.Close(); err != nil {
        return err
    }

    outName := temp.Name()

    fmt.Fprintln(os.Stdout, outName)

    return os.WriteFile(outName, data, 0666)
}

func unaryNodesToString(headRef *UnaryNode) string {
	graphString := "graph TD\n"

	func() {
        from := strNode(headRef.Val)
		headRef = headRef.Next

		if headRef == nil {
			return
		}

        to := strNode(headRef.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)
	}()

    return strings.Trim(graphString, "\n")
}

func binaryNodesToString(headRef *BinaryNode) string {
	graphString := "graph TD\n"

	func() {
        from := strNode(headRef.Val)
		headRef = headRef.Next

		if headRef == nil {
			return
		}

        to := strNode(headRef.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)

		edgeRev := fmt.Sprintf("%s --> %s\n", to, from)
		graphString = fmt.Sprintf("%s%s", graphString, edgeRev)
	}()

    return strings.Trim(graphString, "\n")
}

func multiNodesToString(headRef *MultiNode) string {
	graphString := "graph TD\n"

	headNode := *headRef
	adjRefs := headNode.Adj

	bftQ := make([]*MultiNode, 1+len(adjRefs))
	bftQ[0] = headRef

	for i := 0; i < len(adjRefs); i++ {
		bftQ[i+1] = adjRefs[i]
	}

	fromQIdx := 0
	toQIdx := 1
	for fromQIdx < len(bftQ) {
		if toQIdx >= len(bftQ) {
			fromQIdx++
			if fromQIdx == len(bftQ) {
				continue
				break
			}

			headNode = *bftQ[fromQIdx]

			adjRefs = headNode.Adj
			bftQ = append(bftQ, adjRefs...)
			continue
		}

        from := strNode(headNode.Val)

		toNode := bftQ[toQIdx]

        to := strNode(toNode.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)

		toQIdx++
	}

    return strings.Trim(graphString, "\n")
}

func strNode(val string)string{
        re := regexp.MustCompile(`\s*`)
        fromKey := re.ReplaceAllString(val, "")

		return fmt.Sprintf("%s[%s]", fromKey, val)
}
