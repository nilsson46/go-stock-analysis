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

**Översikt**

Backend-applikationen är en RESTful API-tjänst byggd med Go och Gin-ramverket. Applikationen hanterar aktieinformation och erbjuder funktioner för att lägga till, hämta, uppdatera och ta bort aktier.

**Komponenter**

Gin-ramverket: Används för att skapa RESTful API-endpoints.
PostgreSQL: Används som databas för att lagra aktieinformation.
Redis: Används för sessionhantering. Dock inte implementerade ännu. 
Docker: Används för att containerisera applikationen.
Kubernetes: Används för att orkestrera och hantera containeriserade applikationer.
GitHub Actions: Används för CI/CD-pipeline. 

**Api-Endpoints** 



![endpoints-stock](https://github.com/user-attachments/assets/36437c84-0625-4ac5-b176-eb1b5244a6e4) 





GitHub Actions Workflow
GitHub Actions används för att automatisera byggning, testning och distribution av applikationen. Tanken är att vid varje push till main så körs docker-publish filen och om tester och bygget går igenom så pushar en ny image upp till kubernetes. Kubernetes ligger och kollar efter uppdateringar kontinuerligt så att den tar den nya imagen och uppdaterar. 




