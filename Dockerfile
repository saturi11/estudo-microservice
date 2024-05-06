# Use a imagem oficial do Go
FROM golang:1.17

# Defina o diretório de trabalho no contêiner
WORKDIR /go/src/estudo-microservice

# Copie os arquivos do projeto para o contêiner
COPY . .

# Baixe todas as dependências
RUN go get -d -v ./...

# Compile o aplicativo
RUN go install -v ./...

# Execute o aplicativo
CMD ["estudo-microservice"]