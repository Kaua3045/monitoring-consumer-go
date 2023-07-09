# Monitoring API

## Ferramentas utilizadas
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

## Sobre

Eu criei um projeto para verificar a disponibilidade do seu site através de uma url fornecida, você pode escolher se quer verificar uma única vez, se você quer verificar todos os dias no mesmo horário ou então a cada mês, tendo um painel que mostra a porcentagem de tempo online, podendo filtrar as verificações por um período, você também pode ver qual foi o tempo de resposta da sua aplicação, para cada verificação e o status.

- Porquê decidiu fazer esse projeto?
  - Para aplicar o que eu estudava, aprender mais sobre deploy e aplicar tudo em um projeto grande

- Quais foram os desafios de implementá-lo?
  - Trabalhar pela primeira vez com golang, aprender sobre a linguagem e implementar a conexão com o AWS SQS

- O que eu aprendi com ele?
  - Aprendi sobre golang e como funciona a conexão com banco de dados utilizando golang, conexão com filas da AWS o SQS

## Tabela de conteúdos

- [Features](#features)
- [Requsitos para rodar o projeto](#requisitos)
- [Instruções para executar o projeto](#instruções-para-executar-o-projeto)
- [Contribua com o projeto](#contribuindo-com-o-projeto)
- [Changelog](#changelog)

## Features

Principais features do sistema

1. Se conectar com uma fila do SQS e processar todas as mensagens, pegar a url e pingar ela, após isso salvar no banco de dados

## Requisitos para rodar o projeto

1. Docker e docker-compose
2. GoLang 1.20.3

## Instruções para executar o projeto

1. Baixe a aplicação:
```bash
# Baixando o projeto e acessando o diretorio
git clone https://github.com/Kaua3045/monitoring-consumer-go.git cd monitoring-consumer-go
```

2. Antes de executar a aplicação, você precisa configurar o arquivo .env.example, depois renomeie ele para .env

3. Agora inicie a aplicação:
```bash
# Iniciando a aplicação
go run main.go
```

## Contribuindo com o projeto

Para contribuir com o projeto, veja mais informações em [CONTRIBUTING](doc/CONTRIBUTING.md)

## Changelog

Para ver as últimas alterações do projeto, acesse [AQUI](doc/changelog.md)