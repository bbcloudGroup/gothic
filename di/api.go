package di

func Invoke(invoke interface{}) {
	err := di.Invoke(invoke)
	if err != nil {
		panic(err)
	}
}