package main

import(
	"flag"
	"net/http"
	"log"
	"os/exec"
)

type textStruct struct {
	Text string
}

func commit(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request from: ", req.RemoteAddr)
		cmd := exec.Command("./root/goProjects/onCommit.sh")
		err := cmd.Start()
		if err != nil {
		   log.Fatal(err)
		   }
		log.Println("Waiting for finish")
		err = cmd.Wait()
		log.Printf("command finished with error: %v", err)
	
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var port = flag.String("port", "8800", "The port where the webhook is hosted.")
	flag.Parse()
	http.HandleFunc("/hook", commit)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
