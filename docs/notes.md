## Passo a passo
### /sequence
1. Recebo a lista
2. Coloco cada elemento em uma matriz para facilitar manipulação e diagnóstico
   * Criar uma matriz para cada um dos elementos validos?
3. Verifico cada um dos elementos e os elementos seguintes na diagonal, abaixo  ou a frente.
    * Se proximo elemento for igual, adicionar elemento a uma variável e verificar sequencia.
    * Se comprimento da variável for igual a 4, salvar sequencia em uma nova lista
    * Se lista possuir pelo menos dois elementos, retorne true

## Questões
* Como validar uma sequência que contém 5 letras?
   * Considerar apenas as 4 primeiras e descartar a 5ª?
   * Se uma sequência é validada na diagonal e existe uma sequencia na linha que inicia em uma letra já validada por esse sequência diagonal, deve ser considerada válida?
* Se letra fizer parte de uma sequência válida:
   * Usar map[string]interface{} com as chaves startRow, endRow, startColumn, endColumn e typeSequence para verificar se sequencia de letra já foi encontrada antes.
* No documento não diz se as letras podem ou não se repetir, mas levando em consideração que uma sequência já foi validada, não teria motivo considerar uma nova sequencia de letras apenas por haver uma letra a mais.
   * Decisão: Caso uma sequência já tenha sido validada, considere que aquela letra não pode ser usada para outra sequência. Usar ideia do dicionário.
   * Qual o esforço? Adicionar mais uma parametro para as funções de validação OU verificar o valor antes de enviar para as funções de validação, assim, caso tal elemento já tenha sido validado, não é necessário prosseguir com a validação.
   * Realizado: Caso retornado que uma sequencia foi validada, salvar coordenadas de acordo com o tipo de sequencia encontrada.

## TODO
* Criar .gitignore para remover arquivo do sqlite - FEITO
* Documentar? Não solicitado.