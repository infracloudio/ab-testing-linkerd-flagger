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


## Header based And Weight Based A/B Testing with Flagger
### Delete the HTTPRoute and TrafficSplit object
Delete all HTTPRoute and TrafficSplit objects as flagger creates them internally while testing
  ```
  $ make delete-httpRoute-traffisplit
  ```
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
     $ kubectl apply -f flagger/weight-based.yaml
     ```
  4. Deploy Load generator for matrics analysis
     1. Deploy load generator   
        ```
        $ make deploy-load-generator
        ```
  5. Suppose you have made changes to the code or you have developed a different version of the API.
      1. Build the changes as follows:
         ```
         $ make build-flagger-release
         ```
      3. Deploy new release
         ```
         $ make patch-flagger-release
         ```
  6. Observe the progress of a release
     ```
     $ watch kubectl -n test get canary
     ```
   

    
