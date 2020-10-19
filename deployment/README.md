# Deployment

Here you will find every file needed to deploy a Kubernetes stack.

- [api-deployment.yaml](api-deployment.yaml): Defines the deployment of the API, it has a initContainer that waits for the database service to accept connections and populates the database with some data.
- [api-service.yaml](api-service.yaml): Defines the service required to expose the API as a Varnish endpoint
- [database-deployment.yaml](database-deployment.yaml): Defines the deployment of the Postgres database
- [database-service.yaml](database-service.yaml): Exposes the database so the API can access it
- [dbdata-persistentvolumeclaim.yaml](dbdata-persistentvolumeclaim.yaml): Defines the persistent volume to store database data.
- [secret.yaml](secret.yaml): Defines the required variables such as database user and password. Normally you do not want to commit this file.
- [varnish-configmap.yaml](varnish-configmap.yaml): Holds the configuration file for Varnish
- [varnish-statefulset.yaml](varnish-statefulset.yaml): Defines the varnish deployment.
- [varnish-service.yaml](varnish-service.yaml): Sets Varnish as a Load Balancer in order to access the API
- [kustomization.yaml](kustomization.yaml): [Kustomize](https://kustomize.io/) config file
  
## Requirements

In order to get Varnish working, you have to set up a `ServiceAccount`, just copy this commands:

```bash
$ kubectl create serviceaccount kube-httpcache
$ kubectl apply -f https://raw.githubusercontent.com/mittwald/kube-httpcache/master/deploy/kubernetes/rbac.yaml
$ kubectl create rolebinding kube-httpcache --clusterrole=kube-httpcache --serviceaccount=default:kube-httpcache --user=kube-httpcache
```

## Deploy

To deployt the stack you can use kustomize with:

```bash
kustomize build deployment | kubectl apply -f -
```
