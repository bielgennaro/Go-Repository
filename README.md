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
