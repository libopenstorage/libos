#!/bin/bash
fail() {
  echo "$1"
  exit 1
}

# pre-generated token added to secret
token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAb3BlbnN0b3JhZ2UuaW8iLCJleHAiOjE3NTMxNDA0NDcsImdyb3VwcyI6WyIqIl0sImlhdCI6MTU5NTQ2MDQ0NywiaXNzIjoib3BlbnN0b3JhZ2UuaW8iLCJuYW1lIjoidXNlciIsInJvbGVzIjpbInN5c3RlbS51c2VyIl0sInN1YiI6InVzZXJAb3BlbnN0b3JhZ2UuaW8ifQ.41yebvGhSUlks4_perFh0sORmGnpulEML-7plFa0WWE
kubectl create secret generic token-secret --from-literal=auth-token=$token
kubectl get secret token-secret || fail "failed to create token secrets"

# Create Storage Class
kubectl apply -f demo/e2e/storageclass.yaml || fail "failed to apply storageclass yaml"
kubectl get storageclass openstorage-sc || fail "failed to create storageclass"

# Create PVC
kubectl apply -f demo/e2e/pvc.yaml || fail "failed to apply PVC yaml"
kubectl get pvc openstorage-pvc | grep Bound

# Check PVC with 5 retries
n=0
while true; do
  kubectl get pvc openstorage-pvc | grep Bound 
  if [ $? -eq 0 ]; then
    break
	elif [ $n -gt 5 ]; then
    kubectl describe pvc openstorage-pvc
    fail "PVC not bound after 5 retries"
  else
    echo "PVC not bound"
    kubectl get pvc openstorage-pvc
    n=$((n+1))
    sleep 5  
  fi
done

# Use PVC with MySQL deployment
kubectl apply -f demo/e2e/mysql.yaml || fail "failed to apply deployment yaml"
kubectl get deployment mysql || fail "failed to create deployment"
