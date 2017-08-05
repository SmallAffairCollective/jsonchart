run: install cleanup
	docker run -d -p 6379:6379 --name redis redis:alpine
	docker run --rm -it --link redis:redis jsonchart

dev: install
	docker run -it -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --entrypoint=/bin/sh jsonchart

install: depends
	docker build -t jsonchart .

cleanup: depends
	docker rm -f redis || true

depends:
	docker -v
