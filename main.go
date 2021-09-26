// Copyright(C) 2021. All Rights Reserved.
// Author: lirubing996@gmail.com
// Date: 2021/9/18

// Package main
package main

import (
	"github.com/lirubing996/design-patterns/creational/abstract-factory"
	"github.com/lirubing996/design-patterns/creational/builder"
	"github.com/lirubing996/design-patterns/creational/factory-method"
	"github.com/lirubing996/design-patterns/creational/prototype"
	"github.com/lirubing996/design-patterns/creational/singleton"
	"github.com/lirubing996/design-patterns/tools"
)

func main() {
	factory_method.Run()
	tools.PrintHr2()

	abstract_factory.Run()
	tools.PrintHr2()

	builder.Run()
	tools.PrintHr2()

	prototype.Run()
	tools.PrintHr2()

	singleton.Run()
	tools.PrintHr2()
}

