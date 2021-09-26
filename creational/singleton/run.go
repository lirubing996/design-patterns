// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/26

// Package singleton
package singleton

import "time"

func Run() {
	for i := 0; i < 30; i++ {
		go getInstance(i)
	}
	time.Sleep(time.Second)
}
