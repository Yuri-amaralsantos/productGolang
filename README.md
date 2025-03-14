# Golang Product API

Este projeto é uma API simples em Golang para gerenciar produtos. A API permite visualizar, criar, editar e deletar produtos, cada um contendo um título e uma descrição.

# Tecnologias utilizadas
Golang

# Gin (Framework para rotas)
# PostgreSQL (Opcional para persistência de dados)

# Instalação

## Clone o repositório:
git clone https://github.com/seu-usuario/seu-repositorio.git

## Navegue até o diretório do projeto:
cd seu-repositorio

## Instale as dependências:
go mod tidy

## Execute o projeto:
go run main.go

# Rotas
## Listar produtos
GET /products
Retorna todos os produtos cadastrados.

## Criar produto
POST /products

Corpo da requisição (JSON):
{
  "title": "Nome do produto",
  "description": "Descrição do produto"
}

## Editar produto
PUT /products/:id

Corpo da requisição (JSON):
{
  "title": "Novo título",
  "description": "Nova descrição"
}

## Deletar produto 
DELETE /products/:id
Remove um produto pelo ID.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto é licenciado sob a licença MIT.

## Autor
Yuri amaral santos - Desenvolvedor 

