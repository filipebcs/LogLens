# Relatório de Análise de Logs

**Arquivo analisado:** `samples/sample.log`

**Data da análise:** 2026-01-22 18:51:40

**Total de eventos analisados:** 5

## Resumo Executivo

Foram identificados **4 achados relevantes**, conforme detalhado a seguir.

## Achados

### Achado 1

- **Regra:** `excessive_errors`
- **Descrição:** Muitos erros detectados
- **Severidade:** high
- **Tipo:** Agregado
- **Nível:** ERROR
- **Contagem:** 4
- **Limite:** 2

### Achado 2

- **Regra:** `burst_errors`
- **Descrição:** Explosão de erros em curto intervalo
- **Severidade:** high
- **Tipo:** Agregado
- **Nível:** ERROR
- **Contagem:** 4
- **Limite:** 3

### Achado 3

- **Regra:** `connection_failure`
- **Descrição:** Falha de conexão detectada
- **Severidade:** high
- **Tipo:** Evidência pontual
- **Nível:** ERROR
- **Linha:** 1
- **Mensagem:** `failed to connect`

### Achado 4

- **Regra:** `connection_failure`
- **Descrição:** Falha de conexão detectada
- **Severidade:** high
- **Tipo:** Evidência pontual
- **Nível:** ERROR
- **Linha:** 2
- **Mensagem:** `timeout`

