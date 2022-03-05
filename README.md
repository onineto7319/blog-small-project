# blog-small-project
---

Blog-Backend-Simple-Porject 是使用 Golang + MySQL 建立的簡易文章及標籤的後端專案，部署於Heroku，以Restful API來進行資料的互動需要。

---
**Features**
- JWT授權
- 使用Git Action實踐自動化測試(CI)
- 整合Heroku 實踐自動化部屬(CD)

---
**Development environment**

- [Golang 1.15+](https://go.dev/)
- [MySQL](https://www.mysql.com/)

---

****DB Structure****

![ERD](./assets/ERD.png)
---

****API Document****

[PostMan](https://www.postman.com/onineto7319/workspace/blog-small-project/request/5284931-703f8f85-6898-407d-baf1-2c1d52b4c7ef)

---

****Project Layout****
```
.
├── configs              
├── docs               
├── global               
├── migrations           
├── internal             
│   ├── router            
│   ├── middleware       
│   ├── service             
│   ├── model           
│   ├── dao           
│   ├── dto           
└── pkg                  
    ├── app        
    ├── database        
    ├── errcode              
    ├── setting      
    └── util
```

---
**Installing**

1. Clone專案
```
$ git clone https://github.com/onineto7319/blog-small-project.git
```

2. 下載依賴的Package
```
$ cd blog-small-project
$ go mod download
```

3. 變數設定更改(./configs/config.yaml)
```
Server:
  HttpPort: 8080
  ReadTimeout: 60
  WriteTImeout: 60
DatabaseMysql:
  DBType: mysql
  Username: root
  Password: 1234
  Host: 127.0.0.1:3306
  DBName: blog_service
  ...
``` 
4. 資料庫設定
新建資料庫名稱為 config.yaml檔內的DBName
5. 啟動專案
```
$ go run main.go
```

