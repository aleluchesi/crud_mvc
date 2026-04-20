# Documento de Especificação de Software (SSD) - CRUD MVC em Go

## 1. Introdução

Este documento descreve a especificação do software para um sistema de gerenciamento de usuários implementado em Go utilizando a arquitetura MVC (Model-View-Controller). O sistema fornece operações CRUD (Create, Read, Update, Delete) para usuários via uma API REST.

### 1.1 Objetivo
O objetivo do software é fornecer uma API RESTful para gerenciar usuários, incluindo criação, consulta, atualização e exclusão de registros de usuários.

### 1.2 Escopo
O sistema inclui:
- Endpoints REST para operações CRUD em usuários
- Validação de dados de entrada
- Tratamento de erros padronizado
- Configuração via variáveis de ambiente

## 2. Requisitos Funcionais

### 2.1 Gerenciamento de Usuários
- **RF1**: Criar um novo usuário
  - Entrada: Nome, email, senha, idade
  - Validações: Nome (2-10 caracteres), email válido, senha (mín. 6 caracteres com caracteres especiais), idade (1-120)
- **RF2**: Buscar usuário por ID
  - Entrada: ID do usuário
  - Saída: Dados do usuário (ID, nome, email, idade)
- **RF3**: Buscar usuário por email
  - Entrada: Email do usuário
  - Saída: Dados do usuário
- **RF4**: Atualizar usuário
  - Entrada: ID do usuário e dados a atualizar
- **RF5**: Excluir usuário
  - Entrada: ID do usuário

## 3. Requisitos Não Funcionais

### 3.1 Performance
- O sistema deve responder às requisições em menos de 500ms para operações simples.

### 3.2 Segurança
- Validação de entrada para prevenir injeção de dados.
- Senhas devem conter caracteres especiais.

### 3.3 Usabilidade
- API RESTful com códigos de status HTTP apropriados.
- Mensagens de erro em português.

### 3.4 Portabilidade
- Executável em sistemas Linux com Go instalado.

## 4. Arquitetura

### 4.1 Padrão MVC
- **Model**: Estruturas de dados para requisições e respostas (`request.UserRequest`, `response.UserResponse`)
- **View**: Não aplicável (API REST)
- **Controller**: Lógica de negócio em `controller/user/`

### 4.2 Tecnologias
- Linguagem: Go 1.24.4
- Framework Web: Gin
- Validação: go-playground/validator
- Configuração: godotenv

## 5. Modelos de Dados

### 5.1 UserRequest
```go
type UserRequest struct {
    Name     string `json:"name" binding:"required,min=2,max=10"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,containsany=@#$%"`
    Age      int8   `json:"age" binding:"required,min=1,max=120"`
}
```

### 5.2 UserResponse
```go
type UserResponse struct {
    Id    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int8   `json:"age"`
}
```

## 6. APIs

### 6.1 Endpoints
- `GET /getUserById/:userId` - Buscar usuário por ID
- `GET /getUserByEmail/:userEmail` - Buscar usuário por email
- `POST /createUser` - Criar usuário
- `PUT /updateUser/:userId` - Atualizar usuário
- `DELETE /deleteUser/:userId` - Excluir usuário

### 6.2 Códigos de Status
- 200: Sucesso
- 400: Bad Request (dados inválidos)
- 500: Internal Server Error

## 7. Tratamento de Erros

O sistema utiliza uma estrutura padronizada para erros REST:
```go
type RestErr struct {
    Message string   `json:"message"`
    Code    int      `json:"code"`
    Summary string   `json:"resume"`
    Causes  []Causes `json:"causes,omitempty"`
}
```

## 8. Configuração

- Porta do servidor definida pela variável de ambiente `PORT_ROUTE`
- Arquivo `.env` para configurações

## 9. Dependências

- github.com/gin-gonic/gin
- github.com/joho/godotenv
- github.com/go-playground/validator/v10

## 10. Conclusão

Este documento fornece uma visão geral completa do sistema CRUD MVC em Go. Para implementações futuras, considere adicionar persistência de dados (banco de dados) e autenticação/autorização.