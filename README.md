# Kubernetes Commands Cheat Sheet

## 1. Create Kubernetes Components from a YAML File
To create resources in Kubernetes using a YAML file, use:
```sh
kubectl apply -f <file-name>.yaml
```

## 2. Delete Kubernetes Components from a YAML File
To remove resources defined in a YAML file:
```sh
kubectl delete -f <file-name>.yaml
```

## 3. Show Kubernetes Components
To list specific Kubernetes resources such as Pods, Jobs, Services, PVCs, ConfigMaps, or Secrets:
```sh
kubectl get <pods|job|svc|pvc|configmap|secret>
```
Example:
```sh
kubectl get pods  # List all Pods
kubectl get svc   # List all Services
```

## 4. Show Logs of a Kubernetes Component
To view logs of a specific component:
```sh
kubectl logs <component-name>
```
Example:
```sh
kubectl logs my-pod
```

## 5. Access Kubernetes Components from Local Machine
To access a running Kubernetes component (Pod or Service) from your local machine:
```sh
kubectl port-forward <component-name> <local-port>:<component-port>
```
Example:
```sh
kubectl port-forward pod/my-app 8080:80
kubectl port-forward svc/my-service 5000:5000
```
Now, you can access the service at `http://localhost:8080`.

## 6. Generate Encrypt Values for secret.yaml in Base64
To store sensitive data securely in Kubernetes secrets, you need to encode values in Base64 before adding them to secret.yaml.
```sh
echo -n "your-secret-value" | base64
```