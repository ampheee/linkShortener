# Информация о решении:

Тестовое задание Ozon-fintech. GRPC-реализация. 
Сервис запускается с помощью Makefile и docker-compose.
Способ хранения определяется Makefile, либо вручную. Выбор - PostgreSQL/Redis

<br>[Старт.](#старт)
<br>[Алгоритм шифрования и почему он работает.](#алгоритм-шифрования-и-почему-он-работает)
<br>[Проверка решения.](#проверка-решения)

## Старт:
    make app-on-redis # старт с Redis в качестве хранилища.
    make app-on-postgres # старт с PostgreSQL в качестве хранилища.
    docker-compose run app -storage=PostgreSQL/Redis # старт командой run.

## Алгоритм шифрования и почему он работает.
 <b>Основная идея разработанного алгоритма заключается именно в том, чтобы получать только один и единственный 
 шифр для каждой оригинальной ссылки и при этом было доступным для понимания. 
 Из-за этого мой выбор пал не на хеш от времени или фантомных переменных, а на 
 sha256 и модифированный алгоритм EncodeBase62/64 для мощности алфавита 63, расписанного в тз. 
 <br>WorkFlow: usecase -link-> HashLink(link string) -n-> EncodeBase63(n uint64) <-> back | 
 <br>Алгоритм получает на вход строку link, хеширует ее с помощью sha256 и парсит хеш в *big.Int, откуда далее
преобразуется в uint64 и попадает в EncodeBase63 где и сокращает строку.</br></b>

## Проверка решения:

  Все эндпоинты находятся на `localhost:8080`
  <br> Эндпоинты данной реализации:
  <br> `/create` с передачей JSON.
  <br> `/{abbreviatedLink}` с передачей сокращенной ссылки формата `rus.tam/HgJ46nyw8E`


### <i> С помощью CLI </i>: 

![screenPostCurl](images/postCurl.png)
<br>`- Запрос POST командой curl`

![screenGetCurl](images/getCurl.png)
<br>`- Запрос GET командой`

### <i> C помощью Postman/Insomnia/etc.. </i>:

![screenPostPostman](images/postPostman.png)
<br>`- Запрос POST c Postman`

![sceenGetPostman](images/getPostman.png)
<br>`- Запрос GET c Postman`
