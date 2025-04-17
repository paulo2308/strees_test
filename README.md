# 🧪 LoadTester CLI

Ferramenta de linha de comando (CLI), escrita em Go, para realizar testes de carga em serviços web. Permite definir a URL alvo, número de requisições e chamadas simultâneas, e exibe um relatório ao final da execução.

---

## 🚀 Funcionalidades

- Teste de carga com configuração personalizada
- Controle de número total de requisições
- Controle de nível de concorrência
- Relatório de:
  - Tempo total da execução
  - Total de requisições
  - Requisições bem-sucedidas (HTTP 200)
  - Outros códigos de status HTTP
  - Número de erros

---

## ⚙️ Como usar

### ✅ Rodando localmente

```bash
go run main.go --url=https://www.google.com --requests=100 --concurrency=10
```

### ✅ Rodando com docker

```bash
docker build -t loadtester .
docker run loadtester \
  --url=http://example.com \
  --requests=100 \
  --concurrency=10
```
