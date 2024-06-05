package main

func main() {

	//store, err := NewPostgresStore()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if err := store.Init(); err != nil {
	//	log.Fatal(err)
	//}
	server := NewAPIServer(":4000")
	server.Run()

	//fmt.Printf("%+v\n", store)
}
