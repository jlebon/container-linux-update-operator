#!/bin/bash

echo "## Creating the name space..."
oc create -f examples/namespace.yaml
echo "## Adding the cluster role..."
oc create -f examples/cluster-role.yaml
echo "## Adding role bindings..."
oc create -f examples/cluster-role-binding.yaml
echo "## Adding the operator..."
oc create -f examples/update-operator.yaml
echo "## Adding the agent..."
oc create -f examples/update-agent.yaml
sleep 5
echo "## Agent DS"
oc get ds --namespace=reboot-coordinator
sleep 5
echo "## reboot-coordinator"
oc get all --namespace=reboot-coordinator
