// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package builder
package builder

type iBuilder interface {
	setHouseType()
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}