# gostrap

simple script to bootstrap new folder structure for simple go backend

## Usage

1. **Clone the Repository**: First, clone the repository to your local machine.

2. **Build the Go Program**: Build the program to create an executable binary. Navigate to the directory containing your `main.go` file and run:

    ```sh
    go build -o gostrap
    ```

3. **Move the Executable to a Directory in Your PATH**: Move the built executable to a directory that is included in your system's `PATH`.
    ```sh
    sudo mv gostrap /usr/local/bin/
    ```

4. **Run Your Program**: Now you should be able to run your program from anywhere you want to initialize a project using the command:

    ```sh
    gostrap
    ```