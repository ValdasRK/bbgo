package types

import (
	"math/rand"
	"testing"

	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/stretchr/testify/assert"
)

func TestRBTree_InsertAndDelete(t *testing.T) {
	tree := NewRBTree()
	node := tree.Rightmost()
	assert.Nil(t, node)

	tree.Insert(fixedpoint.NewFromInt(10), 10)
	tree.Insert(fixedpoint.NewFromInt(9), 9)
	tree.Insert(fixedpoint.NewFromInt(12), 12)
	tree.Insert(fixedpoint.NewFromInt(11), 11)
	tree.Insert(fixedpoint.NewFromInt(13), 13)

	node = tree.Rightmost()
	assert.Equal(t, fixedpoint.NewFromInt(13), node.key)
	assert.Equal(t, fixedpoint.Value(13), node.value)

	ok := tree.Delete(fixedpoint.NewFromInt(12))
	assert.True(t, ok, "should delete the node successfully")
}

func TestRBTree_Rightmost(t *testing.T) {
	tree := NewRBTree()
	node := tree.Rightmost()
	assert.Nil(t, node, "should be nil")

	tree.Insert(10, 10)
	node = tree.Rightmost()
	assert.Equal(t, fixedpoint.Value(10), node.key)
	assert.Equal(t, fixedpoint.Value(10), node.value)

	tree.Insert(12, 12)
	tree.Insert(9, 9)
	node = tree.Rightmost()
	assert.Equal(t, fixedpoint.Value(12), node.key)
}

func TestRBTree_RandomInsertSearchAndDelete(t *testing.T) {
	var keys []fixedpoint.Value

	tree := NewRBTree()
	for i := 1; i < 100; i++ {
		v := fixedpoint.NewFromFloat(rand.Float64()*100 + 1.0)
		keys = append(keys, v)
		tree.Insert(v, v)
	}

	for _, key := range keys {
		node := tree.Search(key)
		assert.NotNil(t, node)

		ok := tree.Delete(key)
		assert.True(t, ok, "should find and delete the node")
	}
}

func TestRBTree_CopyInorder(t *testing.T) {
	tree := NewRBTree()
	for i := 1.0; i < 10.0; i += 1.0 {
		tree.Insert(fixedpoint.NewFromFloat(i*100.0), fixedpoint.NewFromFloat(i))
	}

	newTree := tree.CopyInorder(3)
	assert.Equal(t, 3, newTree.Size())

	newTree.Print()

	node1 := newTree.Search(fixedpoint.NewFromFloat(100.0))
	assert.NotNil(t, node1)

	node2 := newTree.Search(fixedpoint.NewFromFloat(200.0))
	assert.NotNil(t, node2)

	node3 := newTree.Search(fixedpoint.NewFromFloat(300.0))
	assert.NotNil(t, node3)

	node4 := newTree.Search(fixedpoint.NewFromFloat(400.0))
	assert.Nil(t, node4)
}

func TestTree_Copy(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(fixedpoint.NewFromFloat(3000.0), fixedpoint.NewFromFloat(1.0))
	assert.NotNil(t, tree.Root)

	tree.Insert(fixedpoint.NewFromFloat(4000.0), fixedpoint.NewFromFloat(2.0))
	tree.Insert(fixedpoint.NewFromFloat(2000.0), fixedpoint.NewFromFloat(3.0))

	newTree := tree.Copy()
	node1 := newTree.Search(fixedpoint.NewFromFloat(2000.0))
	assert.NotNil(t, node1)
	assert.Equal(t, fixedpoint.NewFromFloat(2000.0), node1.key)
	assert.Equal(t, fixedpoint.NewFromFloat(3.0), node1.value)

	node2 := newTree.Search(fixedpoint.NewFromFloat(3000.0))
	assert.NotNil(t, node2)
	assert.Equal(t, fixedpoint.NewFromFloat(3000.0), node2.key)
	assert.Equal(t, fixedpoint.NewFromFloat(1.0), node2.value)

	node3 := newTree.Search(fixedpoint.NewFromFloat(4000.0))
	assert.NotNil(t, node3)
	assert.Equal(t, fixedpoint.NewFromFloat(4000.0), node3.key)
	assert.Equal(t, fixedpoint.NewFromFloat(2.0), node3.value)
}

func TestTree(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(fixedpoint.NewFromFloat(3000.0), fixedpoint.NewFromFloat(10.0))
	assert.NotNil(t, tree.Root)

	tree.Insert(fixedpoint.NewFromFloat(4000.0), fixedpoint.NewFromFloat(10.0))
	tree.Insert(fixedpoint.NewFromFloat(2000.0), fixedpoint.NewFromFloat(10.0))

	// root is always black
	assert.Equal(t, fixedpoint.NewFromFloat(3000.0), tree.Root.key)
	assert.Equal(t, Black, tree.Root.color)

	assert.Equal(t, fixedpoint.NewFromFloat(2000.0), tree.Root.left.key)
	assert.Equal(t, Red, tree.Root.left.color)

	assert.Equal(t, fixedpoint.NewFromFloat(4000.0), tree.Root.right.key)
	assert.Equal(t, Red, tree.Root.right.color)

	// should rotate
	tree.Insert(fixedpoint.NewFromFloat(1500.0), fixedpoint.NewFromFloat(10.0))
	tree.Insert(fixedpoint.NewFromFloat(1000.0), fixedpoint.NewFromFloat(10.0))

	deleted := tree.Delete(fixedpoint.NewFromFloat(1000.0))
	assert.True(t, deleted)

	deleted = tree.Delete(fixedpoint.NewFromFloat(1500.0))
	assert.True(t, deleted)

}
