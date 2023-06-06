build-a:
	docker build -t backend-a -f Dockerfile.backend-a .

build-b:
	docker build -t backend-b -f Dockerfile.backend-b .

build-forwarder:
	docker build -t forwarder -f Dockerfile.forwarder .

build-all: build-a build-b build-forwarder


load-kind:
	kind load docker-image backend-a
	kind load docker-image backend-b
	kind load docker-image forwarder

deploy-all:
	kubectl apply -f deploy/backend-A -f deploy/backend-B -f deploy/forwarder

configure-httpRoute:
	kubectl apply -f linkerd/httpRoute.yaml
configure-traffic-split:
	kubectl apply -f linkerd/traffic-split.yaml
