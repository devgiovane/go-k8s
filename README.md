```bash
kind create cluster --config=kind.yaml
```

```bash
kubectl get nodes
kubectl get pods 
kubectl get replicasets
kubectl get services
```

kubectl config get-clusters
kubectl config use-context <name>

```bash
kubectl apply -f config/pod.yaml
kubectl apply -f config/deployment.yaml 
kubectl apply -f config/service.yaml 
```

```bash
kubectl port-forward pod/goserver 8000:80
kubectl port-forward svc/goserver-service 8000:80
```