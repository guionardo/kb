---
title: Gerador de caça-palavras
tags:
    - python
    - jogo
---

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
        for index, palavra in enumerate(self.palavras):
            horizontal = index % 2 == 0
            if not self.try_put_word(palavra, horizontal):
                self.try_put_word(palavra, not horizontal)

    def can_put_word_at(self, word: str, x: int, y: int, horizontal: bool) -> bool:
        if (horizontal and x+len(word) >= self.width) or (not horizontal and y+len(word) >= self.height):
            return False
        for index, c in enumerate(list(word)):
            e = self.get_char(
                x+index, y) if horizontal else self.get_char(x, y+index)
            if e not in [' ', c]:
                return False
        return True

    def try_put_word(self, word: str, horizontal: bool) -> bool:
        cols = list(range(self.width-len(word)))
        random.shuffle(cols)
        rows = list(range(self.height-len(word)))
        random.shuffle(rows)
        for col, row in itertools.product(cols, rows):
            if self.can_put_word_at(word, col, row, horizontal):
                self.put_word_at(word, col, row, horizontal)
                return True
        return False

    def fill_empty_chars(self):
        for index, c in enumerate(self.chars):
            if c == ' ':
                self.chars[index] = random.choice(string.ascii_uppercase)

    def put_word_at(self, word: str, x: int, y: int, horizontal: bool):
        for index, c in enumerate(list(word)):
            if horizontal:
                self.set_char(x+index, y, c)
            else:
                self.set_char(x, y+index, c)

    def set_char(self, x: int, y: int, c: str):
        self.chars[y*self.width+x] = c[0] if c else ' '

    def get_char(self, x: int, y: int) -> str:
        return self.chars[y*self.width+x]

    def get_row(self, y: int) -> str:
        return ''.join([self.get_char(x, y) for x in range(self.width)])

    def __str__(self) -> str:
        return '\n'.join([self.get_row(y) for y in range(self.height)])
```

## Uso

```python
from caca_palavra.tabuleiro import Tabuleiro
 
tabuleiro = Tabuleiro(
            ['nome', 'endereço', 'bairro', 'cidade', 'estado', 'país', 'casa', 'camelo', 'amarelo', 'bacana', 'guionardo'])
self.assertEqual(15, tabuleiro.width)
self.assertEqual(15, tabuleiro.height)
tabuleiro.fill_empty_chars()
print(tabuleiro)
```

Saída:

```
LUJQTCASAPHQKSE
DYENDEREÇOJRYFS
BVNAFQBACANAWTF
ANGHYDECIDADEBC
ILUXYHCAMELONYM
RPILISSCXJMSMZC
RRODPPSRMJBICSY
OPNAXAYAMARELON
MPADTÍMQELXUVZD
NXRWISGJSXNWSHB
OUDSLGXMTBQJZON
MSOYIBIQALSGNNX
ENPFKOODDGLVCCP
EOQZCEVBOHBPNNZ
FOUOASQDALQNTWW
```
