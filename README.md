# k8-native-inside-cluster

## Extra manual Activity

This pod won't have privilege to list out all the pods and deployments as its using default service account. We have to create a role and rolebinding to
assign the role to this default service account. Below are the set of commands that I ran locally in minikube to make it work.

```ssh
 PS C:\AC\go-ws\k8-native-inside-cluster> kubectl create role temp --resource pods,deployments --verb list
 role.rbac.authorization.k8s.io/temp created
 PS C:\AC\go-ws\k8-native-inside-cluster> kubectl create rolebinding temp --role temp --serviceaccount default:default
 rolebinding.rbac.authorization.k8s.io/temp created

```