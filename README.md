# learning
a collection of miscellaneous software files for learning and training purposes. 

- 1. Initialize the program
    - create a go.mod file
        - terminal commands: 
            1. 'go mod init <name of mod>'
            2. 'go mod tidy' (this will create and add the go.mod and go.sum file(s))
            3. 'go mod verify'
    - utilize a docker file
        - terminal commands: 
            - build cmd: docker build -t <image-name> .
            - run image: docker run -p 8080:8080 -it <image-name>
        