// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package prototype
package prototype

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tmpChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tmpChildren = append(tmpChildren, copy)
	}
	cloneFolder.children = tmpChildren
	return cloneFolder
}
