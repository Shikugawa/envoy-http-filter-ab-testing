# A/B Testing with Envoy HTTP Filter

### Endpoints
all of endpoints support POST only

- `/login`
```
curl -H 'Content-Type:application/json' -d '{"username":"Taro","password":"nyanpass"}' http://localhost:8080/login
```

- `/welcome`
```
curl -H 'Content-Type:application/json' -H 'x-user-id-mod:0' -d '{"session_id":"6dfabb1f-6e2b-498f-ae59-9665728319be"}' http://localhost:8080/welcome
```


