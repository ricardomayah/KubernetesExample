# KubernetesExample
-----------------Pasos para el despliegue del backend y el front end del app application Example--------------

1. Seguir los pasos del archivo Readme en la carpeta backend

2. Si se encuentra en windows correr el comando 
   kubectl port-forward svc/service-backend-k8s-hands-on 7000:8080

3. Si se encuentra en linux cambiar la url del index.html
    var url = "http://localhost:7000";
    por
    var url = "http://service-backend-k8s-hands-on:8080";

(Nota este cambio se da debido a que el request lo hace el navegador, el servicio que se levanta en el 
backend es de tipo clusterip y solo es reconocido dentro del cluster. Al estar en windows no podemos acceder directamente a la IP o 
al dns service-backend-k8s-hands-on por lo que se crea un portforward para ello)

4. Seguir los pasos del Readme en la carpeta frontend

5. Verificar la IP de la maquina con ipconfig 

6. Verificar el puerto asignado al nodeport con el comando

    kubectl get svc

7. Acceder: ejemplo http://192.168.0.111:30389/
