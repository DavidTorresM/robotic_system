# Sistema para Concursos de Robótica

Este proyecto es un backend para la gestión de concursos de robótica, desarrollado en Go, con autenticación, gestión de equipos, robots, categorías y rondas de competición.


## Move the fingers

---

## Requisitos

- Docker y Docker Compose
- AWS CLI configurado (para despliegue en AWS)
- Make (en Windows puedes usar Git Bash o WSL)

---

## Variables de entorno

Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido:

```env
IP_SERVER=0.0.0.0
PORT_SERVER=3000

POSTGRES_PASSWORD=mysecretpassword
DATABASE_DSN="postgres://postgres:mysecretpassword@db:5432/postgres?sslmode=disable"

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SENDER_EMAIL=tu_email@gmail.com
SENDER_PASSWORD=tu_contraseña

IP_PUBLICA_SERVER=localhost
```

---

## Desarrollo local

### Lanzar la base de datos

```sh
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

### Construir la imagen

```sh
docker build -t project/robotic_system .
```

### Lanzar el contenedor

```sh
docker run -e IP_SERVER=0.0.0.0 -e PORT_SERVER=3000 -e DATABASE_DSN="postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" -p 3000:3000 project/robotic_system
```

### O usando Docker Compose

```sh
docker compose up --build
```

---

## Despliegue en AWS usando Makefile

El proyecto incluye un `Makefile` para facilitar el despliegue en AWS (ECR y ECS).  
Asegúrate de tener configuradas tus credenciales de AWS y el AWS CLI instalado.

### Comandos principales

- **Crear repositorio ECR y cluster ECS, subir imagen:**
  ```sh
  make deploy_aws
  ```
  Esto ejecuta los pasos:
  1. Crea el repositorio ECR (si no existe)
  2. Crea el cluster ECS (si no existe)
  3. Sube la imagen Docker a ECR

- **Construir la imagen Docker localmente:**
  ```sh
  make build
  ```

- **Ejecutar el contenedor localmente:**
  ```sh
  make run
  ```

- **Eliminar recursos en AWS (ECR y ECS):**
  ```sh
  make aws_prune
  ```

- **Limpiar contenedores e imágenes locales:**
  ```sh
  make prune
  ```

---

## Endpoints principales

Consulta la colección Postman incluida (`robotic_system_postman.json`) para ejemplos de uso de la API.

---

## Notas

- Para producción, recuerda cambiar las contraseñas y datos sensibles.
- Si usas Gmail para SMTP, puede que necesites una contraseña de aplicación.
- El despliegue en ECS requiere definir tareas y servicios adicionales en AWS.

---

¿Dudas o sugerencias? ¡Abre un issue o contacta al autor!













