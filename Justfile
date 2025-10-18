build:
	cd src/sampleapp && go build -o ../../bin/sampleapp .

run: build
	./bin/sampleapp apiserver

run-port port:
    ./bin/sampleapp apiserver --port {{port}}

test:
	cd src/sampleapp && go test ./...

tidy:
	cd src/sampleapp && go mod tidy
