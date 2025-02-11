package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

var createdDirs = make([]string, 0)
var createdFiles = make([]string, 0)

func getPackageName() string {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nOperation cancelled.")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter package name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading package name: %v", err)
	}

	if strings.TrimSpace(name) == "" {
		log.Fatalf("Please enter a valid package name")
	}

	return name[:len(name)-1]
}

func initGoMod(pkgName string) error {
	cmd := exec.Command("go", "mod", "init", pkgName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createDirs() error {
	dirs := []string{
		"cmd/app",
		"internal/app/handlers",
		"internal/app/models",
		"internal/app/repositories",
		"internal/app/services",
		"internal/pkg/db",
		"internal/pkg/logger",
		"internal/pkg/config",
		"pkg/shared",
		"api/v1",
		"scripts",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("Error creating directory %s: %v", dir, err)
			return err
		}
		createdDirs = append(createdDirs, dir)
	}

	return nil
}

func createFiles() error {
	files := map[string]string{
		".env":       "",
		".gitignore": "*.log\n*.tmp\n*.env",
		"main.go":    "package main\n\nfunc main() {\n\t// TODO: Implement main function\n}\n",
	}

	for file, content := range files {
		f, err := os.Create(file)
		if err != nil {
			log.Printf("Error creating file %s: %v", file, err)
			return err
		}
		createdFiles = append(createdFiles, file)

		if content != "" {
			_, _ = io.WriteString(f, content)
			f.Close()
		}
	}

	return nil
}

func cleanUp() {
	var err error
	msg := "Error during cleanup, please cleanup manually"
	for _, dir := range createdDirs {
		err = os.RemoveAll(dir)
		if err != nil {
			log.Fatalf("%s", msg)
			os.Exit(1)
		}
	}

	for _, file := range createdFiles {
		err = os.Remove(file)
		if err != nil {
			log.Fatalf("%s", msg)
			os.Exit(1)
		}
	}
}

func main() {
	pkgName := getPackageName()
	if err := initGoMod(pkgName); err != nil {
		log.Fatalf("Error initializing go module: %v", err)
		return
	}

	if err := createDirs(); err != nil {
		cleanUp()
	}

	if err := createFiles(); err != nil {
		cleanUp()
	}

	log.Println("Project structure created successfully.")
}
