# Use a imagem base do Go 1.17
FROM golang:1.17

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src


# Defina o comando padrão para executar o aplicativo
CMD ["tail" ,"-f","/dev/null"]