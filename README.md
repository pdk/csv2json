# csv2json

Bare-bones converter from CSV to JSON. Reads stdin and writes stdout. Data is
completely loaded and then written after, so everything must fit in RAM.

ex:

    csv2json < my-data.csv > my-data.json
