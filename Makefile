build:
	docker build -t custom_golang:1 .
run:
	docker run --rm -it -p 4000:4000 custom_golang:1
