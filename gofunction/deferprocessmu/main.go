package deferprocessmu

import "sync"

func doSomething() error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	r1, err := OpenResource1()
	if err != nil {
		return err
	}
	defer r1.Close()

	r2, err := OpenResource2()
	if err != nil {
		return err
	}
	defer r2.Close()

	r3, err := OpenResource3()
	if err != nil {
		return err
	}
	defer r3.Close()

	// 使用r1，r2, r3
	return doWithResources()
}
