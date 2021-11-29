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
	icon := flag.Bool("default-icon", false, "whether the default icon should be included")
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
			// line 0 is the summary
			summary = line
		} else if lineNum == 1 {
			// line 1 is the seperator
			if line != "---" {
				log.Fatalln("Missing three-dashes (---) separator on line 1 between summary and description")
			}
			lineNum++
			continue
		} else {
			// lines 2-end are the description
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
