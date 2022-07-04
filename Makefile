BINARY_NAME=onkyo-remote

build:
	cd frontend; npm run build
	cp -ra frontend/public/* static/
	mv static/index.prod.html static/index.html
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin *.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux *.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows.exe *.go

deps:
	go mod tidy
	cd frontend; npm install

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows.exe
	rm -r static/*
	touch static/index.html
	rm -rf frontend/public/build
	rm -rf frontend/node_modules

dev:
	rm -rf static/*
	echo "Start the frontend using <code>npm run dev</code> from within the <code>frontend/</code> directory" > static/index.html
	go run *.go -config config.yaml
