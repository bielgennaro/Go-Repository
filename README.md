### Progress tracking: Go

 - 15 janeiro de 2024:
	 - Aprendendo sobre imports e como funciona a formatação da linguagem, funções variádicas e blank identifier e alguns conceitos básicos
	 - Declarações de variáveis de tipagem automática
   - Override de valores em variáveis
   - Marmota (:=) só funciona dentro de code blocks (funções)
   - Keywords: Palavras chaves que não podem ser usadas como identificadores
   - Operadores e operando

 - 16 de janeiro de 2024 :
    - Introdução ao workspace com multiplus módulos (go.work)
    - Para criar uma variável fora de escopo é necessário usar a palavra "var" e isso é chamado de "package level" (nivel de package pq package está em um escopo global)
    - Quando uma variável é settada fora do code block ela não pode ser alterada dentro de um code block, isso significa que se vc criou um int fora do code block, ela só podera ser settada como int dentro do code block
    - Valor zero:
      - ints: 0
      - floats: 0.0
      - booleans: false
      - strings: ""
      - pointers, functions, interfaces, slices, channels, maps: nil
    - Use a marmota (:=) sempre que possível
    - Use "var" para package level scope
    ## Pacote fmt
    - Raw strings: Não há interpretação, será mostrado no output da forma com que é colocada em código
    - Intepreted strings: É interpretada podendo conter quebras de linhas (\n) e tabs (\t) dentro da string
    - Grupo 1#: Print -> standart out
      - func Print(a ...interface{}) (n int, err error)
      - func Println(a ...interface{}) (n int, err, error)t //Adiciona uma linha nova no final
      - func Printf()
    - Grupo 2#: Print -> string, pode ser usado como variável
      - func Sprint(a ...interface{}) string
      - func Sprintf(format string, a ...interface{}) string
      - func Sprintln(a ...interface{}) string //Adiciona uma linha nova no final
    - Grupo 3#: Print -> file, writer interface, e.g. arquivo ou resposta do servidor
      - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
      - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
      - func Fprintln(w io.Writer, a ...interface{}) //Adiciona uma linha nova no final

 - 17 de janeiro de 2024 :
    - Criando tipos, tipos são fixos. Uma vez declarados eles não podem ser alterados
    - Type hotdog int -> var b hotdog (main hotdog)
    - Uma variável do tipo hotdog pode ser atribuida com o valor de uma variável tipo int, mesmo que seja subjacente de hotdog
    - Conversão de tipos e não coerção!!
    - Em Go só existe CONVERSÃO de tipos!!
    - a = int(b)
    - ref/spec#Conversions
    - Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
    - Exercicio 1 nivel 1 RESOLVIDO ✔️

    - Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
    - Exercicio 2 nivel 1 RESOLVIDO ✔️

    - Utilizando a solução do exercício anterior:
    1. Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        2. Demonstre a variável "s".
    - Exercicio 2 parte 2 nivel 1 RESOLVIDO ✔️

    1. Crie um tipo. O tipo subjacente deve ser int.
      - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
    2. Na função main:
      1. Demonstre o valor da variável "x"
      2. Demonstre o tipo da variável "x"
      3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
    - Agora já somos quase ninjas nível 1!
    - Exercicio 3 nivel 1 RESOLVIDO ✔️

    - Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"
     - Exercicio 3 parte 2 nivel 1 RESOLVIDO ✔️

     - Agora era pra ter o link de um forms de prova, mas o curso é de 3 anos atrás e o link foi excluído... isso significa que eu passei hehe

  - 18 de janeiro de 2024 :
    - Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
    - Ná prática:
      - Zero value
      - Atribuindo um valor
      - Bool como resultado de operadores racionais