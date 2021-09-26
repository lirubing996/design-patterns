// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/26

// Package singleton
package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance(id int) *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.", id)
			singleInstance = &single{}
		} else {
			fmt.Println("Creating single already created.", id)
		}
	} else {
		fmt.Println("The single already created.", id)
	}
	return singleInstance
}