## summary

test [jizhang-serverSide](https://github.com/guangtuan/jizhang) in BDD way

## framework

[godog](https://github.com/cucumber/godog)

## store

put everything into memory, use `k(string(id))-v(json(body))` pairs

## env

```shell
export JIZHANG_TEST_API_HOST=${host}
godog
```