/*
This script takes markdown input in the following format:
	SUMMARY
	---
	DESCRIPTION

And outputs XML in the following format:
	<component>
	<summary>SUMMARY</summary>
	<description><![CDATA[
	DESCRIPTION
	]]></description>
	<icon type="local">hardcoded_relative_path</icon>
	</component>
*/

package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	// appstreamExt the suffix of output appstream file
	appstreamExt = ".metainfo.xml"
	// iconPath is the path to icon, relative path to the output appstream file
	iconPath = "edgex-snap-icon.png"
)

func main() {
	input := flag.String("input", "", "path to md file")
	flag.Parse()

	if *input == "" {
		fmt.Println("Missing input.\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var component struct {
		XMLName     xml.Name `xml:"component"`
		Summary     string   `xml:"summary"`
		Description struct {
			// use CDATA to make the value human-readable but
			// not interpreted as XML markup
			Body string `xml:",cdata"`
		} `xml:"description"`
		Icon struct {
			Path string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"icon"`
	}

	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Printf("Reading %s", file.Name())

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
			// lines 2 to EOF are the description
			description += line + "\n"
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Summary:\n%s\n\n", summary)
	// fmt.Printf("Description:\n%s\n\n", description)

	if strings.TrimSpace(summary) == "" {
		log.Fatal("Failed to extract summary")
	}
	if strings.TrimSpace(description) == "" {
		log.Fatal("Failed to extract description")
	}

	component.Summary = summary
	component.Description.Body = "\n" + description
	component.Icon.Type = "local"
	component.Icon.Path = iconPath

	output, err := xml.MarshalIndent(&component, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Output:\n%s\n\n", output)

	filename := strings.TrimSuffix(file.Name(), ".md") + appstreamExt
	log.Printf("Writing output to %s", filename)
	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
