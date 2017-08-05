run: install
	docker run -it jsonchart

dev: install
	docker run -it -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --entrypoint=/bin/sh jsonchart

install: depends
	docker build -t jsonchart .

depends:
	docker -v
