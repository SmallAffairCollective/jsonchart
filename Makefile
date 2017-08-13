run: clean install examples generator redis
	docker run --rm -d -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --link jsonchart-examples:jsonchart-examples --link redis:redis --link genit:genit --name jsonchart -p 8080:8080 jsonchart
	docker logs -f jsonchart

local: clean go examples generator redis
	go build
	@echo ""
	@echo "some examples to try out:"
	@echo ""
	@echo "\t./jsonchart http://localhost:8000/genit 1 2 localhost"
	@echo "\t./jsonchart http://localhost:8888/example2.json 2 2 localhost"
	@echo ""

examples: clean install
	docker run -d -p 8888:8000 --name jsonchart-examples jsonchart-examples

generator: clean install
	docker run -d -p 8000:80 --name genit genit

dev: install
	docker run -rm -it -v $$(pwd):/go/src/github.com/SmallAffairCollective/jsonchart --entrypoint=/bin/sh jsonchart

redis: clean
	docker run -d -p 6379:6379 --name redis redis:alpine

install: depends
	docker build -t jsonchart .
	cd examples/generator && docker build -t genit .
	cd examples && docker build -t jsonchart-examples .

clean: depends
	docker rm -f redis || true
	docker rm -f genit || true
	docker rm -f jsonchart-examples || true
	docker rm -f jsonchart || true

depends:
	docker -v

go:
	go version
