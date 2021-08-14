package tree

import (
	"fmt"
	"io/fs"
	"os"
)

type TreeNode struct {
	path        string
	depth       int
	fileInfo    fs.FileInfo
	isLast      bool
	outerPrefix string
	innerPrefix string
}

func NewTreeNode(path string, depth int, fileInfo fs.FileInfo, isLast bool, outerPrefix string) *TreeNode {
	return &TreeNode{
		path:        path,
		depth:       depth,
		fileInfo:    fileInfo,
		isLast:      isLast,
		outerPrefix: outerPrefix,
		innerPrefix: getInnerPrefix(isLast),
	}
}

func (treeNode *TreeNode) Depth() int {
	return treeNode.depth
}

func (treeNode *TreeNode) Path() string {
	return treeNode.path
}

func (treeNode *TreeNode) Display(out *os.File) {
	fmt.Fprintln(out, treeNode.outerPrefix+treeNode.innerPrefix+treeNode.fileInfo.Name())
}

func GetOuterPrefix(treeNode *TreeNode) string {
	if treeNode.isLast {
		return treeNode.outerPrefix + "\t"
	}
	return treeNode.outerPrefix + "│   "
}

func getInnerPrefix(isLast bool) string {
	if isLast {
		return "└──"
	} else {
		return "├──"
	}
}
