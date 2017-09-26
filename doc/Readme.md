# HTTP Header
`application/json`
```java
{
    "client-type": "PRINTER", // PRINTER, APP
    "id": "unique id", // PRINTER Mac address, User token ...
}
```

# User

`GET user/printers`
Return a list of available printers
```javascript
{
    printers: [
        {
            id: "123",
            name: "Vincent's printer"
        }
    ]
}
```

`POST user/printers/{{printerid}}/messages`
send new message to the printer
```javascript
{
    text: "Hello world"
}
```

`GET user/printers/{{printerid}}/messages`
get all the messages sent to the printer
```javascript
{
    messages: [
        {
            sent: 12345056, // UNIX timestamp
            printed: 1234056 // UNIX timestamp (can be null if not printed)
            text: "Hello world",
        }
    ]
}
```

# Printer
`POST /printer`
Register new printer
```javascript
{
    uuid: "123", // Printer unique identifier (MAC Address)
    name: "Vincent's printer" // Printer name
}
```

`GET /printer/messages`
Get all the messages that needs to be printed
```javascript
{
    messages: [
        {
            id: "123",
            text: "Hello world"
        }
    ]
}
```

`PUT /printer/messages/{{messageid}}/printed`
Mark message as printed
