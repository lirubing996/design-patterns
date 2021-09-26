// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package builder
package builder

type normalBuilder struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setHouseType() {
	b.houseType = "normal"
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		houseType:  b.houseType,
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
