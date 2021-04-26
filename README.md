# gomaintest

Пример тестирования web сервера, как чёрного ящика с подсчётом покрытия. В качестве простейшего теста выступает curl.

Запуск:
```
$ ./run_tests.sh
wait-for-it: waiting 15 seconds for 127.0.0.1:8090
wait-for-it: 127.0.0.1:8090 is available after 1 seconds
ok      github.com/Grey-Fox/gomaintest  2.010s  coverage: 80.0% of statements in ./...
github.com/Grey-Fox/gomaintest/main.go:12:      hello           100.0%
github.com/Grey-Fox/gomaintest/main.go:16:      headers         0.0%
github.com/Grey-Fox/gomaintest/main.go:29:      main            100.0%
total:                                          (statements)    80.0%
```
