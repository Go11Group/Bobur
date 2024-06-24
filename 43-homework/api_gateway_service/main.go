package main

import "model/api"

func main() {
	panic(api.Routes().ListenAndServe())
}
