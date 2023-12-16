Варианты запуска:
```
make run        локально, in-memory хранилище
make dbrun      docker-compose сервер + хранилище, хранилище postgresql
make memorun    docker для сервера, хранилище in-memory
```

Удаление контейнеров:
```
make dbstop     
make memostop 
```

Запуск тестов:
```
make test
```