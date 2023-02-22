# crawler-one

O rastreador deve ser uma ferramenta de linha de comando que aceita um URL inicial e um
diretório de destino. O rastreador fará o download da página no URL, salve-a em
o diretório de destino e prossiga recursivamente para quaisquer links válidos nesta página.
Um link válido é o valor de um atributo href em uma tag <a> que resolve para urls que
são filhos da URL inicial. Por exemplo, determinado URL inicial https://start.url/abc ,
URLs que resolvem para e

são URLs válidos, mas aqueles que resolvem
ou para não são URLs válidos e

deve ser ignorado.

https://start.url/abc/foo

https://start.url/abc/foo/bar
https://another.domain https://start.url/baz

Além disso, o rastreador deve:
• Tratar corretamente ser interrompido por Ctrl-C
• Execute o trabalho em paralelo, quando razoável
• Suporte a funcionalidade de retomada verificando o diretório de destino para
páginas baixadas e pule o download e o processamento onde não for necessário
• Fornecer cobertura de teste de “caminho feliz”