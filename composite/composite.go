package composite

import "fmt"

type Folder struct {
	Component []Component
	name      string
}

func (f *Folder) Search(keyword string) Component {
	fmt.Printf("在%s 文件中寻找名为%s的文件\n", f.name, keyword)
	if f.name == keyword {
		return f
	}
	for _, composite := range f.Component {
		res := composite.Search(keyword)
		if res != nil {
			return res
		}
	}
	return nil

}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) Add(c Component) {
	f.Component = append(f.Component, c)
}

type Component interface {
	Search(string) Component
	GetName() string
}

type File struct {
	Name string
}

func (f *File) Search(keyword string) Component {
	fmt.Printf("在 %s 中找 %s\n", keyword, f.Name)
	if f.Name == keyword {
		fmt.Println("找到了")
		return f
	}
	return nil
}

func (f *File) GetName() string {
	return f.Name
}
