# Passo a passo
## /sequence
1. Recebo a lista
2. Coloco cada elemento em uma matriz para facilitar manipulação e diagnóstico
   * Criar uma matriz para cada um dos elementos validos?
3. Verifico cada um dos elementos e os elementos seguintes na diagonal, abaixo  ou a frente.
    * Se proximo elemento for igual, adicionar elemento a uma variável e verificar sequencia.
    * Se comprimento da variável for igual a 4, salvar sequencia em uma nova lista
    * Se lista possuir pelo menos dois elementos, retorne true