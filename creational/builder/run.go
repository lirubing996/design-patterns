// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/25

// Package builder
package builder

func Run() {
	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	normalHouse.toString()

	iglooBuilder := getBuilder("igloo")
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	iglooHouse.toString()
}

