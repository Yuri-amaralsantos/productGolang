# Product API

Este projeto consiste em uma API REST em Go para gerenciamento de produtos, utilizando o framework Gorilla Mux e o ORM GORM para integração com um banco de dados PostgreSQL. Além disso, há uma interface web básica em HTML e JavaScript para interação com a API.

## Tecnologias Utilizadas
- Go
- Gorilla Mux (roteamento)
- GORM (ORM para PostgreSQL)
- PostgreSQL
- HTML, CSS e JavaScript (frontend básico)

## Funcionalidades
- Criar um produto
- Listar todos os produtos
- Atualizar um produto
- Deletar um produto

## Como Executar o Projeto

### 1. Configurar o Banco de Dados
Certifique-se de ter um banco de dados PostgreSQL rodando. Utilize o seguinte comando SQL para criar o banco de dados necessário:

```sql
CREATE DATABASE productdb;
```

### 2. Configurar as Credenciais do Banco
Edite a variável `dsn` no código para refletir as credenciais corretas do seu banco de dados:

```go
dsn := "host=localhost user=postgres password=admin dbname=productdb port=5432 sslmode=disable"
```

### 3. Instalar Dependências
Instale as dependências do projeto executando:

```sh
go mod tidy
```

### 4. Rodar a Aplicação
Execute o seguinte comando para iniciar o servidor:

```sh
go run main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Endpoints da API

### Criar Produto
**POST** `/products`

```json
{
  "name": "Produto Exemplo",
  "description": "Descrição do produto",
  "price": 99.99
}
```

### Listar Produtos
**GET** `/products`

### Atualizar Produto
**PUT** `/products/{id}`

```json
{
  "name": "Produto Atualizado",
  "description": "Nova descrição",
  "price": 79.99
}
```

### Deletar Produto
**DELETE** `/products/{id}`

## Interface Web
Uma interface básica em HTML e JavaScript está disponível para testar as funcionalidades da API. Basta abrir o arquivo `index.html` no navegador e interagir com os formulários.

## Contribuições
Sinta-se à vontade para contribuir com melhorias para este projeto!

## Licença
Este projeto é distribuído sob a licença MIT.

