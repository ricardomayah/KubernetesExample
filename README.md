![example workflow](https://github.com/ricardomayah/KubernetesExample/actions/workflows/go.yml/badge.svg)
![example workflow](https://github.com/ricardomayah/KubernetesExample/actions/workflows/docker-image.yml/badge.svg)
# KubernetesExample

Se tiene 2 pipelines de ejemplo usando github actions:
1. El primero no utiliza docker y realiza el build y deploy con la configuración para un proyecto en go.
2. El segundo utiliza las imagenes de docker para el push y pull correspondiente de las imagenes.

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



#Este comando crea un contenedor de docker con la imagen de go 
# y se crea un bind mount desde la ruta en donde se corre el comando
# para este ejemplo es: F:\Cursos_siigo\Kubernetes\applicationExample\backend\src
# PWD hace referencia a la ruta antes mencionada y goApp es el directorio en el contenedor
# Nota: este comando correrlo en powershell ya que en git bash da problemas al montar la ruta relativa usando PWD

docker run --rm -dit -w /goApp -p 8716:9090 -v $PWD/:/goApp --name golang golang bash