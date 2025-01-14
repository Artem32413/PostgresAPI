[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&repeat=false&width=435&lines=%D0%9F%D0%A3%D0%9F%D0%9A%D0%9E%D0%92+%D0%90%D0%A0%D0%A2%D0%81%D0%9C+%D0%A2%D0%95%D0%A1%D0%A2%D0%9E%D0%92%D0%9E%D0%95+%D0%97%D0%90%D0%94%D0%90%D0%9D%D0%98%D0%95+%D0%9D%D0%90+GO+GO)](https://git.io/typing-svg)

---

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&repeat=false&width=435&lines=%D0%9E%D0%BF%D0%B8%D1%81%D0%B0%D0%BD%D0%B8%D0%B5+%D0%BF%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D0%B0%3A)](https://git.io/typing-svg)

Этот проект предназначен для демонстрации навыков программирования на языке Go. Он включает в себя различные функции и решения, которые позволяют продемонстрировать знание синтаксиса, структур данных и методов разработки в Go.
REST API сервис для получения, обновления, добавления и удаления сущностей (cars, furniture, flowers) на _PostgreSQL_

Сервис умеет обрабатывать такие завпросы как:

- __Get__

  - Вписываем в Postman > http://localhost:8080/flowers и получаем:
```json
[
    {
        "id": "9",
        "name": "Роза",
        "quantity": 15,
        "price": 200,
        "arrivaldate": "2023-10-01"
    },
    {
        "id": "10",
        "name": "Лилия",
        "quantity": 10,
        "price": 150,
        "arrivaldate": "2023-10-02"
    },
    {
        "id": "11",
        "name": "Тюльпан",
        "quantity": 20,
        "price": 100,
        "arrivaldate": "2023-10-03"
    },
    {
        "id": "12",
        "name": "Гербера",
        "quantity": 25,
        "price": 80,
        "arrivaldate": "2023-10-04"
    }
]
```
---
- __Get id__

  - Вписываем в Postman > http://localhost:8080/flowers/11 и получаем:
```json
[
    {
        "id": "11",
        "name": "Тюльпан",
        "quantity": 20,
        "price": 100,
        "arrivaldate": "2023-10-03"
    }
]
```
---
- __Post__

  - Вписываем в Postman > http://localhost:8080/flowers и строку ниже
```json
[
    {
        "name": "Бегония",
        "quantity": 15,
        "price": 250,
        "arrivaldate": "2023-11-04"
    }
]
```
и получаем:
```json
[
    {
        "id": "13",
        "name": "Бегония",
        "quantity": 15,
        "price": 250,
        "arrivaldate": "2023-11-04"
    }
]
```
---
- __Put__

  - Вписываем в Postman > http://localhost:8080/flowers/12 и строку ниже
```json
{
    "name": "Тюльпан",
    "quantity": 20,
    "price": 100,
    "arrivaldate": "2023-10-03"
}
```
и получаем:
```json
{
    "id": "",
    "name": "Тюльпан",
    "quantity": 20,
    "price": 100,
    "arrivaldate": "2023-10-03"
}
```
---
- __Patch__
  - Вписываем в Postman > http://localhost:8080/flowers/12 и строку ниже
```json
{
    "price": 2000,
    "arrivaldate": "2024-10-06"
}
```
и получаем:

```json
{
    "id": "12",
    "name": "Тюльпан",
    "quantity": 20,
    "price": 2000,
    "arrivaldate": "2024-10-06"
}
```

---
[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&duration=4000&pause=1000&repeat=false&width=550&lines=%D0%9F%D0%BE%D0%B4%D0%B3%D0%BE%D1%82%D0%BE%D0%B2%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D1%8B%D0%B5+%D0%B4%D0%B5%D0%B9%D1%81%D1%82%D0%B2%D0%B8%D1%8F+%D0%B4%D0%BB%D1%8F+%D1%83%D1%81%D0%BF%D0%B5%D1%88%D0%BD%D0%BE%D0%B9+%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D1%8B%3A)](https://git.io/typing-svg) 

  Подготовительные действия Go:
   - Образ с Postgres и самим приложением прописаны в Docker-compose.yml 
     - С помощью команды ``docker-compose up --build app`` можно сбилдить и запустить все приложение
---

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&duration=4000&pause=1000&repeat=false&width=550&lines=%D0%9F%D1%80%D0%B8%D0%BC%D0%B5%D1%80%D1%8B+%D1%81%D1%83%D1%89%D0%BD%D0%BE%D1%81%D1%82%D0%B5%D0%B9+%D0%B8+%D0%B8%D1%85+%D0%BF%D0%BE%D0%BB%D0%B5%D0%B9%3A)](https://git.io/typing-svg)  

```json
    "cars": [
      {
        "id": "",
        "brand": "",
        "model": "",
        "mileage": ,
        "owners": 
      }
    ]

    "furniture": [
      {
        "id": "",
        "name": "",
        "manufacturer": "",
        "height": ,
        "width": ,
        "length": 
      }
    ]
    
    "flowers": [
      {
        "id": "",
        "name": "",
        "quantity": ,
        "price": ,
        "arrival_date": ""
      }
    ]
```