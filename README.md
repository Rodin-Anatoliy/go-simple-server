//task
Написать сервер, дублирующий функционал этого API, c помощью http.Client.
Нужно отдавать общую структуру пользователя c данными профиля, не отображать staticData, password.
Если сумма на аккаунте превышает 50000 скрыть email, username, firstName, lastName, avatar

{
    "id": 28,
    "email": "Garrison.Bechtelar22@yahoo.com",
    "amount": 39912,
    "profile": {
      "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/837.jpg",
      "lastName": "Hettinger",
      "firstName": "Bessie",
      "staticData": "static words from Armenia"
    },
    "password": "ZGe0A6EZIc",
    "username": "Ceasar_Brekke",
    "createdAt": "08/19/2022 00:04:31",
    "createdBy": "2024-09-06 06:57:08.009736738 +0000 UTC m=+649914.512755866"
}


{
    "id": 28,
    "amount": 39912,
    "profile": {},
    "createdAt": "08/19/2022 00:04:31",
    "createdBy": "2024-09-06 06:57:08.009736738 +0000 UTC m=+649914.512755866"
}
