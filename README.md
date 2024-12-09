Scraping could be a option... 

Start to add some manually to the db or in the frontend later on. 
 TODO 

 buildspec.yaml? 

För att köra commando med docker?

För att köra tester så lägger vi till dom i dockerfilen. RUN go test. 

Kubectl commands: 
kubectl --kubeconfig=.\kubeconfig.yaml apply -f .\namespace.yaml 
Om det har gjort ändringar som behöver uppdateras i namespacet. 

kubectl --kubeconfig=.\kubeconfig.yaml apply -f .\pvc.yaml 

Möjlighet till databas i en Persistent volume. 


kubectl --kubeconfig=.\stock-analysis-kubeconfig.yaml apply -f .\postgresql.yaml 

Ingress controller IP: 172.233.35.72
kubectl --kubeconfig=./kubeconfig.yaml  apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.1/deploy/static/provider/cloud/deploy.yaml
kubectl --kubeconfig=./stock-analysis-kubeconfig.yaml  get all -n ingress-nginx 

run it 
docker build -t simonnilsson584/backend:latest .
docker-compose up



docker build -t simonnilsson584/backend:latest .