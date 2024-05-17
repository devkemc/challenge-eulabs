# API de CRUD para Produtos

## Descrição

Esta API foi desenvolvida para um desafio técnico para uma vaga de desenvolvedor backend em Golang. Ela oferece
funcionalidades completas de CRUD (Create, Read, Update, Delete) para gerenciar produtos em um sistema de comércio
eletrônico.

## Tecnologias Utilizadas

Backend: Go (Golang) com o framework Echo e Sqlc
Banco de Dados: MySQL
Documentação: Swagger
Gerenciamento de Dependências: Go Modules
Arquitetura
A API segue uma arquitetura limpa e modular, separando as responsabilidades de negócios, acesso ao banco de dados e
lógica de apresentação.

## Como Executar Localmente

- ### Pré-requisitos
  - Go instalado
  - MySQL instalado ou em docker
- ### Instalação
 - Clone o repositório:
 - git clone https://github.com/devkemc/challenge-eulabs.git
 - Navegue até o diretório do projeto:
 - cd challenge-eulabs/pkg/config/
 - Crie o arquivo config.yaml com as configurações necessárias, incluindo detalhes do banco de dados e qualquer outra
  configuração específica do ambiente.
- Volte para a raiz do projeto:
- cd../../..
- Execute a API:
- go run cmd/api/main.go
 ### Por padrão, a API será acessível em http://localhost. A porta pode variar dependendo da configuração no arquivo config.yaml.

## Documentação

A documentação completa da API, incluindo endpoints, parâmetros esperados e respostas, pode ser encontrada no Swagger
UI. Para acessar a documentação:

Vá para http://localhost:<sua_porta>/swagger/index.html.
Explore os endpoints disponíveis e seus detalhes.

## Postman
 Link para coleções do Postman
 https://elements.getpostman.com/redirect?entityId=24547735-97f4e851-8b68-46f5-81ec-eb3b10afa77a&entityType=collection