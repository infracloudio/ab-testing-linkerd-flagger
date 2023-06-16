# Build and Deploying App Locally

## Build image of all components 
For kind cluster build all the required images.
  ```
  $ make build-all
  $ make load-kind
  ```
For minikube cluster build the required images as follows:
  ```
  $ eval $(minikube docker-env)
  $ make build-all
  ```

## Deploy Application on the kubernetes cluster
Deploy book-svc, book-svc-v1 and Forwarder
  ```
  $ make deploy-all
  ```