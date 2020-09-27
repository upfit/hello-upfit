build:
	go build -o hello-upfit

deploy:
	docker build . -t upfit/hello-upfit
	docker push upfit/hello-upfit