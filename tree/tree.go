package tree

import (
	"fmt"
	"io"
	"io/fs"
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

func (treeNode *TreeNode) IsDir() bool {
	return treeNode.fileInfo.IsDir()
}

func (treeNode *TreeNode) Display(out io.Writer, printFiles bool) {
	nodeInfo := treeNode.outerPrefix + treeNode.innerPrefix + treeNode.fileInfo.Name()

	if printFiles && !treeNode.IsDir() {
		size := treeNode.fileInfo.Size()
		var sizeStr string
		if size > 0 {
			sizeStr = fmt.Sprintf("(%vb)", size)
		} else {
			sizeStr = "(empty)"
		}
		nodeInfo = nodeInfo + " " + sizeStr
	}
	fmt.Fprintln(out, nodeInfo)
}

func GetOuterPrefix(treeNode *TreeNode) string {
	if treeNode.isLast {
		return treeNode.outerPrefix + "	"
	}
	return treeNode.outerPrefix + "│" + "	"
}

func getInnerPrefix(isLast bool) string {
	if isLast {
		return "└───"
	} else {
		return "├───"
	}
}
