# golang-supabase-tasklist
O objetivo desse projeto é aprofundar meus conhecimentos na Golang, realizando uma integração básica com o Supabase, a fim de iniciar meus estudos nessa ferramenta.  
Para a arquitetura do projeto será aplicado alguns conceitos de Ports and Adapters.

## Estrutura

### Entidades

#### Creator

| Propriedade | Tipo   | Obrigatório |
|-------------|--------|-------------|
| id          | int32  | Sim         |
| nome        | string | Sim         |
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
| isFinish    | boolean | Sim         |

##### Ações
- Quero poder **cadastrar** uma `Task`
- Quero poder **encontrar** uma `Task` **por Id**
- Quero poder **encontrar** uma `Task` **pela descrição**
- Quero poder **listar** minhas `Tasks`.
  - Quero poder **ordenar** minhas `Tasks` **com base na data**.
  - Quero poder **ordenar** minhas `Tasks` **pelo estado de conclusão**.
