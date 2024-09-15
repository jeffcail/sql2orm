package main

import (
	"fmt"
	"github.com/jeffcail/sql2orm/handler"
	"log"
	"net/http"
)

const BANNER = `
             .__   ________                              
  ___________|  |  \_____  \ ___  ______________  _____  
 /  ___/ ____/  |   /  ____/ \  \/  /  _ \_  __ \/     \ 
 \___ < <_|  |  |__/       \  >    <  <_> )  | \/  Y Y  \
/____  >__   |____/\_______ \/__/\_ \____/|__|  |__|_|  /
     \/   |__|             \/      \/                 \/ 
`

func main() {
	http.HandleFunc("/gen", handler.GenerateHandler)
	http.Handle("/", http.FileServer(http.Dir("dist")))
	fmt.Printf("%s", BANNER)
	fmt.Println("listening on :7892")
	log.Fatal(http.ListenAndServe(":7892", nil))
}
