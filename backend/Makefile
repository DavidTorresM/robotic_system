# Definir el nombre de la imagen
REGION_AWS = us-east-2
REPOSITORY_NAME_ECR = roboticsystembackend
CLUSTER_NAME_ECS = robotic_systembackend_cluster
IMAGE_NAME = davidcorp/roboticsystembackend
CONTAINER_NAME = roboticsystembackend_container
ACCOUNT_ID = $(shell aws sts get-caller-identity --query "Account" --output text)

# Despliega la aplicacion a aws 
deploy_aws:create_ecr create-ecs upload-image 
	echo "Lista para lanzar a AWS"


run_ecs:
	echo "Pendiente lanzar en ECS"

upload-image:
	aws ecr get-login-password --region ${REGION_AWS} | docker login --username AWS --password-stdin ${ACCOUNT_ID}.dkr.ecr.us-east-2.amazonaws.com
	docker tag ${IMAGE_NAME}:latest ${ACCOUNT_ID}.dkr.ecr.us-east-2.amazonaws.com/project/robotic_system:latest
	docker push ${ACCOUNT_ID}.dkr.ecr.us-east-2.amazonaws.com/project/robotic_system:latest

create-ecs:
	aws ecs create-cluster --cluster-name ${CLUSTER_NAME_ECS} --region ${REGION_AWS} --no-cli-pager || echo "*Error al crear el cluster"

create_ecr:
	-aws ecr create-repository --repository-name ${REPOSITORY_NAME_ECR} --region ${REGION_AWS} --no-cli-pager || echo "*Error al crear el repositorio de imagenes"

aws_prune:
	aws ecr delete-repository --repository-name ${REPOSITORY_NAME_ECR} --region ${REGION_AWS} --force || echo "*Error al eliminar el repositorio de imagenes"
	aws ecs delete-cluster --cluster ${CLUSTER_NAME_ECS} --region ${REGION_AWS} --no-cli-pager || echo "*Error al eliminar el cluster"

# Construir la imagen Docker
build:
	docker build -t $(IMAGE_NAME) .

# Ejecutar un contenedor desde la imagen
run: build
	docker run -d --name $(CONTAINER_NAME) -p 3000:3000 $(IMAGE_NAME)


scan: 
	docker scout quickview  $(IMAGE_NAME)

# Ejecutar el contenedor sin construccion 
r: 
	docker run -d --name $(CONTAINER_NAME) -p 3000:3000 $(IMAGE_NAME)

# Detener el contenedor en ejecución
stop:
	docker stop $(CONTAINER_NAME)

# Eliminar el contenedor detenido
rm:
	docker rm $(CONTAINER_NAME)

# Eliminar la imagen
rmi:
	docker rmi $(IMAGE_NAME)

# Limpiar: detener y eliminar el contenedor
clean: stop rm

# Eliminar imagen y contenedor
prune: clean rmi
