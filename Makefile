# Atlas Makefile
imageName = thedevtop/atobs:latest
	@echo 'build: Build exectuable'
	@echo 'clean: Remove exectuable'
	@echo 'docker: Build Docker image'
build:
	@go fmt
	@go build
clean:
	@go clean
docker:build
	@docker build -t $(imageName) .