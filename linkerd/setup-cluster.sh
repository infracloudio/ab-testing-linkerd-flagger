#!/bin/sh

#configur cluster with linkerd
linkerd install --crds | kubectl apply -f -

linkerd install | kubectl apply -f -

linkerd check

#install linkerd-smi extension
curl -sL https://linkerd.github.io/linkerd-smi/install | sh

linkerd smi install | kubectl apply -f -

linkerd smi check

#install linkerd-viz extension
linkerd viz install | kubectl apply -f -

linkerd viz check