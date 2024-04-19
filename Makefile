.PHONY: build clean deploy

build:
	# app/hello
	sh tools/go_build.sh hello

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy function -f $(function) --verbose

deploy_all: clean build
	sls deploy --verbose
