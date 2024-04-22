.PHONY: build clean deploy

build:
	# app/hello
	sh tools/go_build.sh hello

	# app/hello2
	sh tools/go_build.sh hello2

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy function -f $(function) --verbose --stage $(stage)

deploy_all: clean build
	sls deploy --verbose --stage $(stage)
