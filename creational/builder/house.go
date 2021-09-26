// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package builder
package builder

import "fmt"

type house struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

func (h *house) toString() {
	fmt.Printf("%s House Door Type: %s\n", h.houseType, h.doorType)
	fmt.Printf("%s House Window Type: %s\n", h.houseType, h.windowType)
	fmt.Printf("%s House Num Floor is: %d\n", h.houseType, h.floor)
}
