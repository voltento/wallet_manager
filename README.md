Check the `docs` folder
**Start**

`cd .deploy &&  docker-compose build && docker-compose up`

**API**

Add new account 

`curl -XPUT -d'{"id":"f1", "currency": "USD"}' localhost:8080/account_managing/add/`

Get accounts

`curl localhost:8080/browsing/accounts`

Get payments

`curl localhost:8080/browsing/browsing/payments`

Change account balance

`curl -XPUT -d'{"id":"f1", "change_amount": 90}' localhost:8080/payment/change_balance`

Transfer money

`curl -XPUT -d'{"from_account":"James", "to_account": "foo", "change_amount": 90}' localhost:8080/payment/send_money`
