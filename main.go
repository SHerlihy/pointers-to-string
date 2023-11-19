package main

import (
	"errors"
	"fmt"
	"os"
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

func NodeToStr[T Node](headRef *T) (string, error) {
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

	return retStr, err
}

func unaryNodesToString(headRef *UnaryNode) string {
	graphString := "graph TD\n"

	func() {
		from := fmt.Sprintf("%s[%s]", headRef.Val, headRef.Val)
		headRef = headRef.Next

		if headRef == nil {
			return
		}

		to := fmt.Sprintf("%s[%s]", headRef.Val, headRef.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)
	}()

	return graphString
}

func binaryNodesToString(headRef *BinaryNode) string {
	graphString := "graph TD\n"

	func() {
		from := fmt.Sprintf("%s[%s]", headRef.Val, headRef.Val)
		headRef = headRef.Next

		if headRef == nil {
			return
		}

		to := fmt.Sprintf("%s[%s]", headRef.Val, headRef.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)

		edgeRev := fmt.Sprintf("%s --> %s\n", to, from)
		graphString = fmt.Sprintf("%s%s", graphString, edgeRev)
	}()

	return graphString
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

		from := fmt.Sprintf("%s[%s]", headNode.Val, headNode.Val)

		toNode := bftQ[toQIdx]
		to := fmt.Sprintf("%s[%s]", toNode.Val, toNode.Val)

		edgeFwd := fmt.Sprintf("%s --> %s\n", from, to)
		graphString = fmt.Sprintf("%s%s", graphString, edgeFwd)

		toQIdx++
	}

	return graphString
}
