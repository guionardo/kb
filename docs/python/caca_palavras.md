---
title: Gerador de caça-palavras
tags:
    - python
    - jogo
---

## Código

```python
"""
Caça Palavras
"""

import itertools
import random
import string
from typing import List


class Tabuleiro:
    """Tabuleiro de Caça-palavras
    """

    def __init__(self, palavras: List[str], width: int = 0, height: int = 0):
        """Inicializa o tabuleiro com as palavras escolhidas.
        Largura e Altura são automaticamente detectadas"""
        self.palavras = [p.upper().strip() for p in palavras]
        max_length = (
            max((len(p) for p in self.palavras)) + 6
            if palavras
            else max(width, height)
        )
        self.width = max(max_length, width)
        self.height = max(max_length, height)
        random.seed()
        random.shuffle(self.palavras)
        self.chars = [' ']*self.width*self.height
        horizontal = True
        for palavra in self.palavras:
            if self._try_put_word(palavra, horizontal):
                horizontal = not horizontal
            else:
                self._try_put_word(palavra, not horizontal)
        if self.palavras:
            self._fill_empty_chars()

    def _can_put_word_at(self, word: str, x: int, y: int, horizontal: bool) -> bool:
        """Verifica se uma palavra pode ser inserida na linha ou coluna
        em espaços livres ou com caracteres compartilhados com outras 
        palavras"""
        if (horizontal and x+len(word) >= self.width) or (not horizontal and y+len(word) >= self.height):
            return False
        for index, c in enumerate(list(word)):
            e = self._get_char(
                x+index, y) if horizontal else self._get_char(x, y+index)
            if e not in [' ', c]:
                return False
        return True

    def _try_put_word(self, word: str, horizontal: bool) -> bool:
        """Tenta inserir uma palavra no tabuleiro, procurando por uma
        posição disponível"""
        cols = list(range(self.width-len(word)))
        random.shuffle(cols)
        rows = list(range(self.height-len(word)))
        random.shuffle(rows)
        for col, row in itertools.product(cols, rows):
            if self._can_put_word_at(word, col, row, horizontal):
                self._put_word_at(word, col, row, horizontal)
                return True
        return False

    def _fill_empty_chars(self):
        """Preenche os espaços vazios com caracteres aleatórios"""
        for index, c in enumerate(self.chars):
            if c == ' ':
                self.chars[index] = random.choice(string.ascii_uppercase)

    def _put_word_at(self, word: str, x: int, y: int, horizontal: bool):
        """Insere uma palavra na posição informada"""
        for index, c in enumerate(list(word)):
            if horizontal:
                self._set_char(x+index, y, c)
            else:
                self._set_char(x, y+index, c)

    def _set_char(self, x: int, y: int, c: str):
        """Define um caractere na tabela"""
        self.chars[y*self.width+x] = c[0] if c else ' '

    def _get_char(self, x: int, y: int) -> str:
        """Obtém um caractere da tabela"""
        return self.chars[y*self.width+x]

    def _get_row(self, y: int) -> str:
        """Obtém uma linha da tabela"""
        return ''.join([self._get_char(x, y) for x in range(self.width)])

    def __str__(self) -> str:        
        return '\n'.join([self._get_row(y) for y in range(self.height)])

```

## Uso

```python
from caca_palavra.tabuleiro import Tabuleiro
 
tabuleiro = Tabuleiro([
    'nome', 'endereço', 'bairro', 'cidade',
    'estado', 'país', 'casa', 'camelo',
    'amarelo', 'bacana', 'guionardo'
], 20, 20)

print(tabuleiro)
```

## Saída

```text
JUWYHAYNCYAEHERGDQDQ
OVMKRJWGJBLAMARELOKT
WJRPNXOFWLWTZCIDADEY
ESTADOLMZGGVDBWESWQT
ALORBACANATXOKBNLGMC
JJWRDRPCWRPQLKJOOYNT
DKFWCEFAMHOTAVXMOIYL
PNNPFMEMDELWXABEYVYA
MNPAÍSTEPFAJGKYNHPRH
TLRKMZNLKTGUIONARDOJ
TXGYTWOOPFACAPYKYVJU
YEKQCOHIDGERGCBVDLCZ
LRKTTUUBYKNKAFTJGQSS
EIAXTHRHBXDQNLJMXSGQ
VXSPORCWACESHJVGEUTM
WIAWYMKTIARWRSEQXONS
RLGJVFITRSETDUWBSACL
VVBTETWGRAÇXEKHLQHIA
ISVCQYZWOQOXXZUOHAIM
UYIZVVXGKSISTUGAULON
```
