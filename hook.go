package main

import(
	"flag"
	"net/http"
	"log"
	"os/exec"
	"io/ioutil"
)

type textStruct struct {
	Text string
}

func commit(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request from: ", req.RemoteAddr)
		cmd := exec.Command("./onCommit.sh")
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

var commands []string

func main() {
	
	var port = flag.String("port", "8080", "The port where the webhook is hosted.")
	var file = flag.String("file", "commands.txt", "The commands seperated with newlines.")
	flag.Parse()

	cmdData, err := ioutil.ReadFile(*file)
	check(err)

	cmds := string(cmdData)

	command := ""
	
	for i := 0; i < len(cmds); i++ {
		if string(cmds[i]) != "\n" {
			command += string(cmds[i])
		} else {
			commands = append(commands, command)
			command = ""
		}
	}

	http.HandleFunc("/hook", commit)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

	
}
