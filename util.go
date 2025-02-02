package main

func expectNil(err error) {
	if err != nil {
		panic(err)
	}
}
