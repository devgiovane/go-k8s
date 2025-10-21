```bash
kind create cluster --config=kind.yaml
```

List 
```bash
kubectl get nodes
kubectl get pods 
kubectl get replicasets
kubectl get services
```

```bash
kubectl config get-clusters
kubectl config use-context <name>
```

Apply configs
```bash
kubectl apply -f config/pod.yaml
kubectl apply -f config/deployment.yaml 
kubectl apply -f config/service.yaml 
```

Redirect port
```bash
kubectl port-forward pod/goserver 8000:80
kubectl port-forward svc/goserver-service 8000:80
```

```bash
kubectl proxy --port=8080
```