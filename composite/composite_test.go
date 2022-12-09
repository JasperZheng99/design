package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder3 := &Folder{
		name: "Folder3",
	}

	folder1.Add(folder3)
	folder3.Add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)
	component := folder1.Search("File1")
	if component != nil {
		fmt.Println("找到该文件了")
	}
}
