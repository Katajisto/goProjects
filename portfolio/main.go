package portfolio                                                                   
                                                                               
import (                                                                       
       "fmt"
       "net/http"
       "html/template"
       "math/rand"
       "strconv"
)                                                                              

var codeMap map[string]string;

//Pointers to html templates. tKatajisto
var templatePtr *template.Template
var mainTemplatePtr *template.Template

//Slice that contains all of the texts. tKatajisto
var textSlice []text

func handler(w http.ResponseWriter, r *http.Request) {                         
    keys, ok := r.URL.Query()["code"]
    if !ok || len(keys[0]) < 1 {
        fmt.Fprintf(w,"");
        return
    }
    key := keys[0]
    fmt.Fprintf(w, "%s", codeMap[string(key)]);
}                                                                              

func linkhandler(w http.ResponseWriter, r *http.Request) {
    id := rand.Intn(10000);
    keys, ok := r.URL.Query()["link"]
    if !ok || len(keys[0]) < 1 {
       fmt.Fprintf(w,"")
       return
    }
    key := keys[0];
    codeMap[strconv.Itoa(id)] = string(key);
    fmt.Fprintf(w,"%s",strconv.Itoa(id));
}
                                                                           
func mainh(w http.ResponseWriter, r *http.Request) {                           
     fmt.Fprintf(w, "Tuomas Katajisto collection of APIs.");                   
}                                                                              
                                                                               
func init() {
     //Load texts to memory to improve load times. tKatajisto
     textSlice = parseTexts()
	
     //Parse template files into memory. tKatajisto
     templatePtr = template.Must(template.ParseFiles("/root/goProjects/portfolio/seetext.html"))
     mainTemplatePtr = template.Must(template.ParseFiles("/root/goProjects/portfolio/mainpage.html"))
/*	
     http.HandleFunc("/portfolio/text/", servePage)
     http.HandleFunc("/portfolio/", serveHome)
     codeMap = make(map[string]string)
     codeMap["1337"] = "ktj.st";
     http.HandleFunc("/linkshare/addlink/",linkhandler);
     http.HandleFunc("/linkshare/getcode/",handler);                            
     http.Handle("/linkshare/", http.StripPrefix("/linkshare/", http.FileServer(http.Dir("ls"))))
     http.Handle("/tmacs/", http.StripPrefix("/tmacs/", http.FileServer(http.Dir("tmacs"))))
     http.Handle("/timer/",http.StripPrefix("/timer/",http.FileServer(http.Dir("kenttatimer"))))
     http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("devblog"))))
     //http.HandleFunc("/",mainh);                                             
     log.Fatal(http.ListenAndServe(":80",nil));                                
*/
}      
