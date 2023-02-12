package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Resume struct to store the resume data
type Resume struct {
	ContactInformation struct {
		Name    string `yaml:"name"`
		Email   string `yaml:"email"`
		Address string `yaml:"address"`
		Phone   string `yaml:"phone"`
	} `yaml:"contact_information"`
	Education []struct {
		School    string `yaml:"school"`
		Degree    string `yaml:"degree"`
		Subject   string `yaml:"subject"`
		Graduated string `yaml:"graduated"`
	} `yaml:"education"`
	WorkExperience []struct {
		Company     string `yaml:"company"`
		Position    string `yaml:"position"`
		Duration    string `yaml:"duration"`
		Description string `yaml:"description"`
	} `yaml:"work_experience"`
	Skills []string `yaml:"skills"`
}

func main() {
	// Serve the main route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resume := Resume{}

		// Load the resume data from a static YAML file
		yamlFile, err := ioutil.ReadFile("resume.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = yaml.Unmarshal(yamlFile, &resume)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Parse the HTML template file
		tmpl, err := template.ParseFiles("resume.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Render the HTML template with the resume data
		err = tmpl.Execute(w, resume)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	// Serve the static file server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
