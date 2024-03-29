# golang-supabase-tasklist
O objetivo desse projeto é aprofundar meus conhecimentos na Golang, realizando uma integração básica com o Supabase, a fim de iniciar meus estudos nessa ferramenta.  
Para a arquitetura do projeto será aplicado alguns conceitos de Ports and Adapters.

## Estrutura

![structure](docs/structure.png)

### Entidades

#### Creator

| Propriedade | Tipo   | Obrigatório |
|-------------|--------|-------------|
| id          | string | Sim         |
| name        | string | Sim         |
| email       | string | Sim         |


##### Ações
- Quero poder **cadastrar** um `Creator`
- Quero poder **encontrar** um `Creator` **por Id**
- Quero poder **encontrar** uma `Creator` **pela email**

#### Task

| Propriedade | Tipo    | Obrigatório |
|-------------|---------|-------------|
| id          | int32   | Sim         |
| title       | string  | Sim         |
| description | string  | Não         |
| date        | time    | Sim         |
| creator     | Creator | Não         |
| finished    | bool | Sim         |

##### Ações
- Quero poder **cadastrar** uma `Task`
- Quero poder **encontrar** uma `Task` **por Id**
- Quero poder **encontrar** minhas `Tasks` **pelo título**
- Quero poder **listar** todas as minhas `Tasks`.
  - Quero poder **ordenar** minhas `Tasks` **com base na data**.
  - Quero poder **ordenar** minhas `Tasks` **pelo estado de conclusão**.
