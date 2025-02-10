package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func initGoMod(pkgName string) error {
	cmd := exec.Command("go", "mod", "init", pkgName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createDirs() {
	dirs := []string{
		"cmd/app",
		"internal/app/handlers",
		"internal/app/models",
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
			log.Fatalf("Error creating directory %s: %v", dir, err)
		}
	}
}

func createFiles() {
	files := map[string]string{
		".env":       "",
		".gitignore": "*.log\n*.tmp\n*.env",
		"main.go":    "package main\n\nfunc main() {\n\t// TODO: Implement main function\n}\n",
	}

	for file, content := range files {
		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("Error creating file %s: %v", file, err)
		}

		if content != "" {
			_, err = io.WriteString(f, content)
			f.Close()
		}
	}
}

func getPackageName() string {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	nameChan := make(chan string, 1)

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

//program will accept packagename only
//first initgomod
//create dirs
//create files
//can I make it revert to original state if it encounters an error at any point??

func main() {
	pkgName := getPackageName()
	if err := initGoMod(pkgName); err != nil {
		log.Fatalf("Error initializing go module: %v", err)
		return
	}
	createDirs()
	createFiles()
	log.Println("Project structure created successfully.")
}
