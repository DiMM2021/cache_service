# Cache Service

`Cache Service` — это простая библиотека для создания и управления кешем в памяти. Она позволяет устанавливать, получать и удалять значения, используя ключи.

## Установка

Чтобы установить библиотеку, используйте следующую команду:

```bash
go get -u github.com/yourusername/cache_service

## Использование

package main

import (
    "cache_service/cache"
    "fmt"
)

func main() {
    // Создание нового кеша
    cache := cache.New()

    // Установка значения в кеш
    cache.Set("userId", 42)

    // Получение значения из кеша
    userId := cache.Get("userId")
    fmt.Println(userId)  // Output: 42

    // Удаление значения из кеша
    cache.Delete("userId")
    userId = cache.Get("userId")
    fmt.Println(userId)  // Output: <nil>
}

## API

cache.New()
Создает новый экземпляр кеша.

cache.Set(key string, value interface{})
Устанавливает значение в кеш по указанному ключу.
key: строка, представляющая ключ.
value: значение любого типа, которое нужно сохранить.

cache.Get(key string) interface{}
Возвращает значение, сохраненное в кеше по указанному ключу.
key: строка, представляющая ключ.
Возвращает сохраненное значение или nil, если ключ не найден.

cache.Delete(key string)
Удаляет значение из кеша по указанному ключу.
key: строка, представляющая ключ.

## Лицензия
Этот проект распространяется под лицензией MIT.