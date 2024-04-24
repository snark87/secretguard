package secretguard

import "github.com/awnumar/memguard/core"

func safePanic(p any) {
	core.Panic(p)
}

func EnsureSafePanic() {
	if r := recover(); r != nil {
		safePanic(r)
	}
}
