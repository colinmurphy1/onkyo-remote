.PHONY: frontend linux windows darwin
BINARY_NAME=onkyo-remote

linux: frontend
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux *.go

windows: frontend
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows.exe *.go

darwin: frontend
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin *.go

# Compile everything
all: linux windows darwin

# Compile just the frontend
frontend:
	cd frontend; npm run build
	cp -ra frontend/public/* static/
	mv static/index.prod.html static/index.html

# Install all Go and node dependencies
deps:
	go mod tidy
	go mod vendor
	cd frontend; npm install

# Clean up
clean:
	go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows.exe
	rm -r static/*
	touch static/index.html
	rm -rf frontend/public/build
	rm -rf frontend/node_modules

# Start the development backend (needs a config.yaml in this directory to function)
dev:
	rm -rf static/*
	echo "Start the frontend using <code>npm run dev</code> from within the <code>frontend/</code> directory" > static/index.html
	go run *.go -config config.yaml
