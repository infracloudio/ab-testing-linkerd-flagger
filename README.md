# dynamicRoutingApp

## Prerequisites
- Kubernetes kind/minikuke cluster
- Linkerd must be installed and configured with the cluster(https://linkerd.io/2.13/tasks/install/)
- [Linkerd-SMI](https://linkerd.io/2.13/tasks/linkerd-smi/#cli) and [Linkerd-viz](https://linkerd.io/2.13/tasks/troubleshooting/#l5d-viz-ns-exists) extension must be installed.
- We can install Linkerd and required addons on the Kubernetes cluster in one step with
  ```
  $ make setup-cluster-linkerd
  ```

## Create and configure the namespace
Create and configure namespace with linkerd proxy so that Deployments created in the namespace will be automatically configured with linkerd sidecar/proxy
  ```
  kubectl create ns test --dry-run=client -o yaml \
    | linkerd inject - \
    | kubectl apply -f -
  ```


## Build the docker images
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
Deploy backend-a, backend-b and Forwarder
  ```
  $ make deploy-all
  ```

NOTE: We should not enable Dynamic routing and traffic-splitting simultaneously.
## Header based Dynamic Routing
### Enable Dynamic Routing
  - To dynamically route traffic to another version of application i.e backend-b we need configure httpRoute with the following
    ```
    $ make configure-httpRoute
    ```
### Test With Request Header
  - We can dynamically route the traffic to backend-b by using the request header `x-backend: test` 
  - If we want to configure dynamic routing with some different header we can do it by editing the header name and value in the linkerd/httpRoute.yaml file
  - Expose forwarder service
    ```
    $ kubectl port-forward svc/forwarder-service 4000:8080 -n test
    ```
  - Use `x-backend: test` header to control traffic
    ```
    $ curl -sX GET -H 'x-backend: test' localhost:4000/
    ```

## Traffic Splitting
- We can enable traffic split by the ratio 50:50 between the two versions/backend as follows:
  ```
  $ make configure-traffic-split
  ```
- Expose forwarder service
  ```
  $ kubectl port-forward svc/forwarder-service 4000:8080 -n test
  ```
- Test traffic-split with the `curl -sX GET localhost:4000/` such that request will be routed to both version/backend according traffic split ratio.
- We can adjust ratio of traffic by editing the value of weight in the linkerd/traffic-split.yaml file

## Header based And Weight Based A/B Testing with Flagger
### Header based A/B Testing
  1. Install flagger
     ```
     $ make setup-cluster-flagger
     ```
  2. Deploy backend-a and forwarder as follows:
     ```
     $ make deploy-flagger-release
     ```
  3. Apply flagger header based canary object
     ```
     $ kubectl apply -f flagger/header-based.yaml
     ```
  4. Suppose you have made changes to the code or you have developed a different version of the API.
      1. Build the changes as follows:
         ```
         $ make build-flagger-release
         ```
      2. Copy the image tag from the above image build and replace it with image tag in the value field of flagger/flagger-release-patch.json
         ```
         {
            "op": "replace",
            "path": "/spec/template/spec/containers/0/image",
            "value": "backend-a:<your-image-tag>"
         }
         ```
         we are replacing image with new build and version as 'canary' for the deployment backend-a
      3. Deploy new release
         ```
         $ make patch-flagger-release
         ```
  5. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```
### Weight Based A/B Testing
  To implement weight-based A/B testing follow similar steps as header-based A/B testing. But instead of applying a header-based canary, we need to apply a weight-based flagger Canary object as follows     in the third step.
  ```
  $ kubectl apply -f flagger/weight-based.yaml
  ```
