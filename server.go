package main

import (
  "net/http"
  "log"
  "html/template"
  "strings"
  "fmt"
  languagetranslator "github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"

)

type TextBox struct {
Name string
Nametranslang string
}

type DropDown struct {
Name string
Value string
}



type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
  PageTitle        string         
  Answer           string
  TextBoxPV       TextBox

}


var finalstring string
//var texttotranslate string

func editingjson(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}



func main() {

  http.HandleFunc("/", inputfunc)
  http.HandleFunc("/selected", UserSelected)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func translate(ttotranslate string, transcode string) string{
  // Instantiate the Watson Language Translator service
  service, serviceErr := languagetranslator.
    NewLanguageTranslatorV3(&languagetranslator.LanguageTranslatorV3Options{
      URL:       "https://gateway-wdc.watsonplatform.net/language-translator/api", //change this only if needed
      Version:   "2018-11-01",
      IAMApiKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",  //change this only
    })

  // Check successful instantiation
  if serviceErr != nil {
    fmt.Println(serviceErr)
    return ""
  }
  /* TRANSLATE */

textToTranslate := []string{
    "",
    ttotranslate,
  }

  translateOptions := service.NewTranslateOptions(textToTranslate).
    SetModelID(transcode)

  // Call the languageTranslator Translate method
  response, responseErr := service.Translate(translateOptions)

  // Check successful call
  if responseErr != nil {
    panic(responseErr)
  }
//  fmt.Println(response)
   finalstring=editingjson(response.String(), "\"translation\": \"")
   finalstring=finalstring[:len(finalstring)-34]
  fmt.Println(strings.Split(finalstring,"\"")[0])
  return finalstring
}





func inputfunc(w http.ResponseWriter, r *http.Request){
   Title := "TRANSLATION USING WATSON TRANSLATOR API"
  

  MyTextBox := TextBox{"TB", "TB2"}

  MyPageVariables := PageVariables{
    PageTitle: Title,
    TextBoxPV: MyTextBox,
    }

   t, err := template.ParseFiles("select.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
   }
}





func UserSelected(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  // r.Form is now either
  // map[animalselect:[cats]] OR
  // map[animalselect:[dogs]]
 // so get the animal which has been selected
  texttotranslate := r.Form.Get("TB")
  transcode:= r.Form.Get("TB2")

MyTextBox := TextBox{"TB", "TB2"}

  Title := "Translation Achieved"
  MyPageVariables := PageVariables{
    PageTitle: Title,
    Answer : translate(texttotranslate, transcode),
    TextBoxPV: MyTextBox,
    }



 // generate page by passing page variables into template
    t, err := template.ParseFiles("select.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}
