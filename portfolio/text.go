package portfolio

import(
	"io/ioutil"
	"log"
	"bufio"
	"os"
)

type text struct {
	Title string
	Text string
}

func parseTexts() []text {
	var textS []text
	files, err := ioutil.ReadDir("./root/goProjects/portfolio/texts/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		
		//Prevent emacs backups from being read as texts. tKatajisto
		if string(f.Name()[len(f.Name())-1]) == "~" {
			continue
		}
		
		ff, err := os.Open("./root/goProjects/portfolio/texts/"+f.Name())
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(ff)
		content, _ := ioutil.ReadAll(reader)
		newText := text{
			Title: f.Name(),
			Text: string(content),
		}
		textS = append(textS, newText)
	}
	return textS
}
