// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package prototype
package prototype

type inode interface {
	print(string)
	clone() inode
}