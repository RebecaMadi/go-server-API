# go-server-API

* Sobre o projeto

Esta aplicação é minha primeira API utilizando Go, ela serve uma página web e consegue usar os métodos POSt e GET para guardar e exibir dados. O projeto foi feito como um dos meus desafios semanais para aprimorar os meus conhecimentos em Go e desenvolvimento Web. No deafio foi utilizado p pacote gorilla/mux para o melhor gerenciamento de rotas.

* Aprendizados

Técnicamente aprendi um pouco sobe o pacote gorilla/mux, sobre rotas, o pacote net/http e sobre como funcionam as requisições e respostas e maneiras de servir paginas e arquivos.

Os aprendizados que tive com esse desafio vão além do conhecimento tecnico, é um desafio simples, mas inicialmente optei, inconscientemente, pela maneira mais dificil e não prática de ser feita, o que me fez passar mais tempo que o necessário no desafio, entretanto, esse erro me fez visualizar na prática como funcionam as requisições do front com o servidor, após todo o tempo gasto eu descobri, lendo a documentação, que existia uma forma muito mais simples de fazer o que eu tinha feito. Embora eu tenha conseguido tirar algo de produtivo da forma como fiz pela primeira vez eu não acho que seja uma forma sustentavel de programar, então as lições que aprendi foram:
    * Organize os pensamentos em tópicos
    * Veja quais são as funcionalidades do rpograma
    * Pesquise na documentação
    * Só depois programe

* Descrição

O projeto serve paginas html na porta 8000, possui os seguintes diretórios:
    */login : pagina de login
    */users : mostra os usuarios salvos
    *login-submit : envia um novo usuario para o "banco", é assionado quando apertamos o botão
    
*Configurações e instruções
    * O projeto usa o pacote gorilla/mux, caso seja necessário instalá-lo use o comando go get -u github.com/gorilla/mux
    * No termial da pasta utilize o comando go build para gerar o executavel caso ainda não tenha
    * utilize ./go-server-API para executar o servidor
    * acesse a pagina na porta 8000 preferencialmente na pagina /login
