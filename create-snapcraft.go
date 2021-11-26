package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	name := flag.String("name", "", "name of the snap")
	icon := flag.Bool("upload-icon", false, "whether the default icon should be uploaded")
	flag.Parse()

	if *name == "" {
		fmt.Println("Missing input.\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	snapcraft := make(map[string]interface{})

	templateFile, err := ioutil.ReadFile("template.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(templateFile, &snapcraft)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Open("metadata/" + *name + ".md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var description, summary string
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			summary = line
		} else if lineNum == 1 && line == "---" {
			// separator
			continue
		} else {
			description += line + "\n"
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Summary:\n%s\n\n", summary)
	// fmt.Printf("Description:\n%s\n\n", description)

	if !*icon {
		delete(snapcraft, "icon")
	}
	snapcraft["name"] = name
	snapcraft["summary"] = summary
	snapcraft["description"] = description

	output, err := yaml.Marshal(snapcraft)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", output)

	// output = bytes.ReplaceAll(output, []byte("\n"), []byte("  \n"))
	err = os.MkdirAll("snap", 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("snap/snapcraft.yaml", output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
