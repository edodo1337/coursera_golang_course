package main

import (
	"hw1_tree/tree"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

type FileInfoArray []os.FileInfo

func (files FileInfoArray) Len() int      { return len(files) }
func (files FileInfoArray) Swap(i, j int) { files[i], files[j] = files[j], files[i] }
func (files FileInfoArray) Less(i, j int) bool {
	filename1 := files[i].Name()
	filename2 := files[j].Name()

	return filename1 < filename2
}

func excludeFiles(files []os.FileInfo) (result []os.FileInfo) {
	for _, file := range files {
		if file.IsDir() {
			result = append(result, file)
		}
	}
	return
}

func dfs(out io.Writer, path string, fromNode *tree.TreeNode, printFiles bool) error {
	fromNode.Display(out, printFiles)

	if !fromNode.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if !printFiles {
		files = excludeFiles(files)
	}

	outerPrefix := tree.GetOuterPrefix(fromNode)

	for counter, file := range files {
		isLast := counter == len(files)-1

		nodePath := path + "/" + file.Name()
		nodeDepth := fromNode.Depth() + 1
		node := tree.NewTreeNode(nodePath, nodeDepth, file, isLast, outerPrefix)

		dfs(out, node.Path(), node, printFiles)
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	sort.Sort(FileInfoArray(files))

	if !printFiles {
		files = excludeFiles(files)
	}

	for counter, file := range files {
		isLast := counter == len(files)-1

		nodePath := path + "/" + file.Name()
		node := tree.NewTreeNode(nodePath, 0, file, isLast, "")

		if err := dfs(out, nodePath, node, printFiles); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// out := os.Stdout
	out, _ := os.Create("output.txt")
	defer out.Close()

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

}
