# Flash-Cards


Flash-Cards is a project made by Yassa Taiseer that is used to help make quick questions and answers
  - Users signup/login
  - They can then make flash cards which can be later deleted
  - Helps with studying 

### Tech Stack:

  - GoLang(backend)
  - VanillaJs(frontend animations)
  - Html&CSS(frontend)
  - MySQL(Backend)



### Installation

Flash-Cards requires the MySql download [here](https://www.mysql.com/)
It also requires the installation of Golang 1.15+
#### Mac,Linux,Windows:
```sh
go run main.go
```

### Building Database
DlvrMe runs on a MySQL databases
There is a need for two tables Users and Cards
Name the database ```flashcarddb``` on ```3306``` Port
#### User's Tables
The user's table will look like this:
|VALUE| TYPE  |
| ------ | ------ |
| Username | VARCHAR |
| Password | VARCHAR |
| ID_key | AUTO_INCREMENT_KEY |

#### Deliveries Database
| VALUE  | TYPE |
| ------ | ------ |
| Username | VARCHAR |
| Question | VARCHAR |
| Answer | VARCHAR |
| ID_key | AUTO_INCREMENT_KEY |

