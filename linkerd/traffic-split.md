#Trafficsplit Using Linkerd

## Create and configure the namespace
Create and configure namespace with linkerd proxy so that Deployments created in the namespace will be automatically configured with linkerd sidecar/proxy
  ```
  kubectl create ns test --dry-run=client -o yaml \
    | linkerd inject - \
    | kubectl apply -f -
  ```

## Deploy Application on the kubernetes cluster
Deploy book-svc, book-svc-v1 and Forwarder
  ```
  $ make deploy-all
  ```


NOTE: We should not enable Dynamic routing and traffic-splitting simultaneously.
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