Чтобы запустить бэкэнд-сервис в докере, выполните следующие команды

sudo docker build -t biocad .
sudo docker run -d -p 8080:8080 biocad