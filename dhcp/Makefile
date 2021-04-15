IMAGE = "ixothym/dhcp"

.PHONY: build
build:
	@echo Building docker image
	docker buildx build --platform linux/amd64,linux/arm/v7 -t ${IMAGE} . --push

.PHONY: run
run:
	docker run --net=host -it -p 8080:8080 ${IMAGE}

#.PHONY: push
#push:
#	docker push ${IMAGE}
