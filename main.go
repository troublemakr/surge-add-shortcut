package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	path := os.Getenv("SURGE_CONF")
	proxy := os.Getenv("SURGE_PROXY")
	if len(path) == 0 {
		log.Fatalf("Please set env SURGE_CONF to the path of your surge config file.\n")
	}
	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
	} else {
		suffix := os.Args[1]
		domainExp := `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`
		re := regexp.MustCompile(domainExp)
		if !re.Match([]byte(suffix)) {
			log.Fatalf("Error! Invalid DOMAIN-SUFFIX: %s", suffix)
		}
		cmd(path, suffix, proxy)
	}
}

func cmd(path string, suffix string, proxy string) {
	lines, err := readFile(path, suffix, proxy)
	if err != nil {
		log.Fatal(err)
	}
	err = writeFile(path, lines)

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string, newRule string, proxy string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var lines []string
	rd := bufio.NewScanner(file)
	for rd.Scan() {
		line := rd.Text()
		lines = append(lines, line+"\n")
		if strings.HasPrefix(line, "[Rule]") {
			newRule = fmt.Sprintf("DOMAIN-SUFFIX,%s,%s\n", newRule, proxy)
			lines = append(lines, newRule)
		}
	}
	if err = rd.Err(); err != nil {
		return nil, err
	}
	return lines, err
}

func writeFile(path string, lines []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		return err
	}
	for _, line := range lines {
		file.WriteString(line)
	}
	return nil
}
