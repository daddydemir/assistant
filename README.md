# ASSISTANT
------------

![assistant](https://github.com/daddydemir/assistant/assets/42716195/665fa02d-88e1-4b82-bd4e-109ac24c9102)

-----
Asistant projesi içerisinde bir kuyruk ve telegram API'si bulunmakta ve kuyruğa eklenen veriler anlık olarak telegram'dan bildirim gelecek şekilde çalışmaktadır. Şu anda kuyruğu besleyen iki farklı yazılım mevcut. [Crypto](https://github.com/daddydemir/crypto) projesi 
ve [cryptowaves.app](https://www.cryptowaves.app/relative-strength-index/btc) aresinde oluşturulan svg'yi bana göndermek üzere parse eden bir python script'inden oluşuyor, henüz bu script'i githubda paylaşmadım.


![image (1)](https://github.com/daddydemir/assistant/assets/42716195/d88ce514-57f8-4c00-acf1-86baab05b6a4)


### ÖZET
Crypto projesi günde iki kez çalışarak, o gün içerisinde en çok değişime uğrayan coinleri kuyruğa yazar ve assistant projesi ise onun bildirim olarak bana ulaşmasını sağlar. Benzer şekilde 1 saat arayla çalışan bir python scripti [cryptowaves.app](https://www.cryptowaves.app/relative-strength-index/btc)
adresindeki svg'i sisteme kaydeder daha sonra onu png'ye çevirir ve kuyruğa ekler. Bu web sitesinde coinlerin RSI grafiği bulunuyor, yatırım yaparken kullanıcak herhangi grafikten sadece biri.

---

### TEKNOLOJI
----
- ![Go](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white)
- ![RabbitMQ](https://img.shields.io/badge/RabbitMQ-FF6600.svg?style=for-the-badge&logo=RabbitMQ&logoColor=white)
- ![Telegram](https://img.shields.io/badge/Telegram-26A5E4.svg?style=for-the-badge&logo=Telegram&logoColor=white)
- ![Docker](https://img.shields.io/badge/Docker-2496ED.svg?style=for-the-badge&logo=Docker&logoColor=white)

### RUN | DEPLOY
---

Projede RabbitMQ ve Telegram API'i için bir .env dosyası bulunuyor, onları kendinize uygun bir şekilde değiştirdiğinizde docker container oluşturabilirsiniz. Bunun için aşağıdaki komutu kullanın.

```sh

docker build --tag assistant . 

```
> bu komutun sağlıklı bir şekilde çalışabilmesi için terminalde projenin Dockerfile dosyasının bulunduğu dizine gitmeniz ve orada çalıştırmanız gerekiyor.

```sh

docker run assistant

```
bu komutu kullanarak oluşturulan containeri çalıştırabilirsiniz.

Eğer bana projeyle alakalı bir geliştirme veya bug yada herhangi bir sebepten ötürü ulaşmak isterseniz Discord'u kullanabilirsiniz. ( DEMİRON#1218 )
