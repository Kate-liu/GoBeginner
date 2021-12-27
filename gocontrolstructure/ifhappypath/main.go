package ifhappypath

// 伪代码段1：
func doSomething() error {
	if errorCondition1 {
		// some error logic
		// ... ...
		return err1
	}
	// some success logic
	// ... ...
	if errorCondition2 {
		// some error logic
		// ... ...
		return err2
	}
	// some success logic
	// ... ...
	return nil
}

// 伪代码段2：
func doSomething() error {
	if successCondition1 {
		// some success logic
		// ... ...
		if successCondition2 {
			// some success logic
			// ... ...
			return nil
		} else {
			// some error logic
			// ... ...
			return err2
		}
	} else {
		// some error logic
		// ... ...
		return err1
	}
}
