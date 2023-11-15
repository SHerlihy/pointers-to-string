package ptrToStr

import (
	"testing"
)

func TestInitMinQueue(t *testing.T) {
    nextNode := UnaryNode{
        "next val",
        nil,
    }

    unaryNode := UnaryNode{
        "node val",
        &nextNode,
    }

    retStr := NodeToStr[UnaryNode](&unaryNode)

    t.Logf("\n retVal: %v", retStr)
}
