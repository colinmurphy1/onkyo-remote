BINARY_NAME=onkyo-remote

build:
	cd frontend; \
		npm run build; \
		cp -ra public/* ../static/
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin *.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux *.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows.exe *.go

deps:
	go mod tidy
	cd frontend; npm install

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows.exe
	rm -r static/*
	rm -rf frontend/public/build
	rm -rf frontend/node_modules
