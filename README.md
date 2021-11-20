# Learning
a collection of miscellaneous software files for learning and training purposes. 

- 1. Initialize the program
    - create a go.mod file
        - terminal commands: 
            1. 'go mod init <name of mod>'
            2. 'go mod tidy' (this will create and add the go.mod and go.sum file(s))
            3. 'go mod verify'
- 2. Building a docker image
    - utilize a docker file
        - terminal commands: 
            - build cmd: docker build -t <image-name> .
            - run image: docker run -p 8080:8080 -it <image-name>
    - docker commands: 
        - run: docker run -d -p <port>:<port> -it <image-name>
        - build: docker build --build-arg OS=linux --build-arg ARCH=${arch} -t docker-golang_${arch} .
            - replace ${arch} with amd64 as an example
- 3. Building a multi-Arch Docker Image
    - Create a builder using dockers buildX command(s):
        - 'docker buildx create --name <builder-name>'
        - (to use a builder use this command) 'docker buildx use <builder-name>'
        - (verify what builder using) 'docker buildx inspect <builder-name>'
        - (make active builder) 'docker buildx inspect <builder-name> --bootstrap'
        - (create multi-images w/ command) 
            - docker buildx build --platform <platforms-listed>, -t <repository>/<docker-image> . --push
            - docker buildx build --platform linux/arm64,linux/amd64 -t <docker-image> .
                - Example: 
                    <platforms-listed> : linux/arm64, linux/amd64