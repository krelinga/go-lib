package valid

func OrPanic(v Validator) {
	if err := v.Validate(); err != nil {
		panic(err)
	}
}