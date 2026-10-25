// Package greet is a plain sub-package of the root module (case B).
// dors.com/testGoMajor/v2/greet is NOT a module: a proxy must answer 404 for it.
package greet

func Greet(name string) string { return "hi " + name + " from dors.com/testGoMajor/v2/greet" }
