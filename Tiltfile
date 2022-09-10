local_resource('hello', # nome do serviço
  cmd='go build -o bin/main main.go', # setup
  serve_cmd='bin/main', # execução
  deps=['web', 'main.go'] # arquivos ou pastas que devem ser observadas alterações
)
