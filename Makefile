build-a:
	docker build -t book-svc -f Dockerfile.book-svc .

build-a-v1:
	docker build -t book-svc:v1 -f Dockerfile.book-svc-v1 .

build-forwarder:
	docker build -t forwarder -f Dockerfile.forwarder .

build-load-generator:
	docker build -t load-generator -f Dockerfile.load-generator .

build-all: build-a build-a-v1 build-forwarder build-load-generator

build-git-reg-image:
	docker build -t ghcr.io/infracloudio/book-svc:latest -f Dockerfile.book-svc .
	docker build -t ghcr.io/infracloudio/book-svc:v1 -f Dockerfile.book-svc-v1 .
	docker build -t ghcr.io/infracloudio/forwarder:latest -f Dockerfile.forwarder .
	docker build -t ghcr.io/infracloudio/load-generator:latest -f Dockerfile.load-generator .


load-kind:
	kind load docker-image book-svc
	kind load docker-image forwarder
	kind load docker-image load-generator
	kind load docker-image book-svc:v1

delete-httpRoute-traffisplit:
	kubectl delete -f linkerd/httpRoute.yaml
	kubectl delete -f linkerd/traffic-split.yaml

deploy-flagger-release:
	kubectl apply -f deploy/book-svc/
	kubectl apply -f deploy/forwarder/

deploy-all:
	kubectl apply -f deploy/book-svc -f deploy/book-svc-v1 -f deploy/forwarder

deploy-load-generator:
	kubectl apply -f deploy/load-generator

configure-httpRoute:
	kubectl apply -f linkerd/httpRoute.yaml
configure-traffic-split:
	kubectl apply -f linkerd/traffic-split.yaml

setup-cluster-linkerd:
	./linkerd/setup-cluster.sh

setup-cluster-flagger:
	kubectl apply -k github.com/fluxcd/flagger/kustomize/linkerd

patch-flagger-release:
	kubectl patch deployment book-svc -n test --type='json' -p="$$(cat flagger/flagger-release-patch.json)"
