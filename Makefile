run: install clean
	docker run -d --name genit genit
	docker run -d -p 6379:6379 --name redis redis:alpine
	docker run --rm -it -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --link redis:redis --link genit:genit jsonchart

dev: install
	docker run -it -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --entrypoint=/bin/sh jsonchart

install: depends
	docker build -t jsonchart .
	cd examples/generator && docker build -t genit .

clean: depends
	docker rm -f redis || true
	docker rm -f genit || true

depends:
	docker -v
