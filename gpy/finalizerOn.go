//go:build finalizerOn
// +build finalizerOn

package gpy

import (
	"runtime"
)

func setFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, finalizer)
}
