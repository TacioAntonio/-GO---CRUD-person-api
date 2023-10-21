# GO
> https://gobyexample.com/

## Instalando pacotes externos
Depois de criar ou verificar o arquivo go.mod, você pode usar o comando `go get` para obter pacotes. 
Certifique-se de estar no diretório raiz do seu projeto ao executar o comando `go get`.
```sh
$ go mod init nome-do-seu-modulo
```

## Desinstalar pacotes
```sh
$ go get -u -v nome-do-seu-modulo@none
```

## Limpando pacotes que não são utilizados
```sh
$ go mod tidy
```

## Pacotes
- Atualizando arquivos a cada mudança: https://github.com/cosmtrek/air
- API: https://github.com/gorilla/mux