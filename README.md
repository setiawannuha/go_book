## **How to Run via Container**

1. **Ensure Docker is Running**  
  Make sure Docker is installed and running on your machine.

2. **Customize Environment Variables**  
  Adjust the environment variables as needed in the `.env` file or directly in the `docker-compose.yml` file.

3. **Build and Run the Docker Container**  
  To build and run the Docker containers in the background, use:
   ```sh
   docker-compose up --build -d
4. **Stop and Remove the Docker Container and Volume**
  To stop the running containers and remove associated volumes, use:
   ```sh
   docker-compose down -v
5. **See running containers**
  To see currently running containers, use:
   ```sh
   docker ps --all
   ```
   In the status column, if it contains "Up ... seconds" it means the container is running.
6. **Access the Running Container**
  To access the shell of a running container, use:
   ```sh
   docker exec -it CONTAINER_NAME /bin/sh
   ```
   Replace CONTAINER_NAME with the actual name of the container you want to access.

