package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	xshellSessionDIR := flag.String("i", "", "Xshell session directory")
	windTermUserFile := flag.String("o", "user.sessions", "WindTerm user file")
	showHelp := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	if *xshellSessionDIR == "" {
		log.Print("Xshell session directory is required")
		flag.Usage()
		return
	}

	windTerm := make([]Session, 0)
	xshFiles := scanSessionDir(*xshellSessionDIR)
	for _, xshFile := range xshFiles {
		log.Printf("Processing %s", xshFile)
		session, err := handleXshFile(xshFile)
		if err != nil {
			log.Printf("Fail to generate session info for %s: %v", xshFile, err)
		}
		session.UUID = NewUUID()
		session.Label = getLable(xshFile)
		session.Group = getGroup(*xshellSessionDIR, xshFile)
		session.Icon = "session::square-mediumorchid"
		session.Protocol = "SSH"
		windTerm = append(windTerm, session)
	}
	saveToJson(windTerm, *windTermUserFile)

}

type Session struct {
	UUID             string `json:"session.uuid"`
	Group            string `json:"session.group"`
	Label            string `json:"session.label"`
	Port             int    `json:"session.port"`
	Protocol         string `json:"session.protocol"`
	Target           string `json:"session.target"`
	IdentityFilePath string `json:"ssh.identityFilePath"`
	Icon             string `json:"session.icon"`
	AutoLogin        string `json:"session.autoLogin"`
}

func NewUUID() string {
	uuid, _ := uuid.NewRandom()
	return uuid.String()
}

func scanSessionDir(sessionDir string) []string {
	var sessionFiles []string
	err := filepath.Walk(sessionDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".xsh") {
			sessionFiles = append(sessionFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return sessionFiles
}

func handleXshFile(xshFile string) (Session, error) {
	session := Session{}

	bytes, err := os.ReadFile(xshFile)
	if err != nil {
		log.Printf("Fail to read %s: %v", xshFile, err)
		return session, err
	}

	data := strings.Split(string(bytes), "\n")
	var ip, user string
	for _, line := range data {
		r := strings.ReplaceAll(line, " ", "")
		r = strings.ReplaceAll(r, "\u0000", "")
		r = strings.ReplaceAll(r, "\r", "")

		if strings.HasPrefix(r, "Host=") {
			ip = strings.TrimPrefix(r, "Host=")
		}
		if strings.HasPrefix(r, "UserName=") {
			user = strings.TrimPrefix(r, "UserName=")
		}
		if strings.HasPrefix(r, "Port=") {
			port := strings.TrimPrefix(r, "Port=")
			session.Port, _ = strconv.Atoi(port)
		}
		if strings.HasPrefix(r, "UserKey=") {
			session.IdentityFilePath = strings.TrimPrefix(r, "UserKey=")
		}
	}
	session.Target = user + "@" + ip
	return session, nil
}

func getLable(xshFile string) string {
	filename := filepath.Base(xshFile)
	return strings.TrimSuffix(filename, ".xsh")
}

func getGroup(basePath, xshFile string) string {
	group := filepath.Dir(xshFile)
	group = strings.ReplaceAll(group, basePath, "")
	group = strings.ReplaceAll(group, "\\", ">")
	return strings.TrimPrefix(group, ">")
}

func saveToJson(session []Session, targetFile string) {
	f, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("Fail to create %s: %v", targetFile, err)
	}
	defer f.Close()

	bytes, _ := json.MarshalIndent(session, "", "  ")
	f.Write(bytes)
	log.Printf("Saved %s", targetFile)
}
