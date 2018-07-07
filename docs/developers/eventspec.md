# Fn Project Common Events Specification  and CloudEvent protocol extensions: 

*Experimental*  - this document is not yet implemented. 

This document describes the `cloudevents` function protocol, this protocol is proposed as the default and only protocol for sending and receiving data to functions. 


Cloud events is a meta-standard for defining event types exchanged between distributed systems. Each cloud event has a corresponding type, and function containers may receive events of one type or multiple types depending on the implementation 



Fn supports several common event types and cloud event exxtensions that are used by the platform, these are used to indicate specific result such as unhandled errors, or to communicate with platform components like the HTTP gateway. 


* Raw HTTP interactions encapsulated in events 
* Internal errors 

This doc describes tue specifications for those events: 


# A note on response events

While Fn supports both RPC (request-response) and one-way (request-only) interactions with containers - cloud events are primarily concerned with originating events (corresponding to requests). Valid cloud events may be returned by containers, however the platform will override the following source attributes on any events returned by FDK response events 

*  `source` will be replaced with the platform-specified URI of the function being invoked 
* `eventID` will be replaced with an identifier generated by the Fn platform
* `eventTime` will be replaced with the point in time the function returned the event  

# common cloud event extensions: 

The following extension  are added/set on incoming events to functions to communicate metadata about the function environment: 

`"ioFnProjectDeadline":"2018-01-30T16:52:39.786Z`
 
will be added to events routed to functions to indicate the date/time by which the event must be processed by a function 


# Internal FDK errors 

In the case where an FDK or container detects an error and is able to propagate information such as an error message back to the server. 


Errors from FDKs are represented as a `io.fnproject.fnError` event - 

Produced by FDK: 
```
{
    "cloudEventsVersion" : "0.1",
    "eventType" : "io.fnproject.fnError",
    "eventTypeVersion" : "1.0",
    "contentType": "application/json; charset=utf-8",
    "data" : {
       "message":"An error occured in the function",
       "fatal": false 
    }
}
```


Transformed by platform 
```
{
    "cloudEventsVersion" : "0.1",
    "eventType" : "io.fnproject.fnError",
    "eventTypeVersion" : "1.0",
    "source" : "http://f",
    "eventID" : "A234-1234-1234", // generated 
    "eventTime" : "2018-04-05T17:31:00Z",
    "contentType": "application/json; charset=utf-8",
    "data" : {
       "message":"An error occured in the function",
    }
}
```



# HTTP communication 

The HTTP events encapsulate requests from and responses to an  HTTP gateway: 

## HTTP Request 

The HTTP request event with type `io.fnproject.httpRequest` is used by the Fn HTTP trigger gateway to encapsulate a request from the gateway routed to a container. 

This event and others received via a direct HTTP gateway  `ioFnProjectHttpReq` metadata extension that contais teh following attributes: 

`method` : the HTTP method of the incoming request
`headers` : The HTTP headers of the originating request
`requestURL` : the request URL on which the request was received 

```
PUT /t/myApp/myTrigger HTTP/1.1
Host: fnservice.com
Content-Type: application/json
Authorization: Bearer asdlkjasldkjas


{"my": "data"}
```

```
{
    "cloudEventsVersion" : "0.1",
    "eventType" : "io.fnproject.httpRequest",
    "eventTypeVersion" : "1.0",
    "source" : "https://fnservice.com/t/myApp/mytrigger",
    "eventID" : "A234-1234-1234",
    "eventTime" : "2018-04-05T17:31:00Z",
    "extensions" : {
        "ioFnProjectHttpReq" : {
            "method":"PUT",
            "headers": {
                "Content-type" : ["application/json"],
                "Authorization": ["Bearer asdlkjasldkjas"]
             },
             "requestURL" : "https://fnservice.com/t/myApp/mytrigger"
         }
         
    },
    "contentType" : "application/json",
    "data" : {"my": "data"}
}

```


```
GET /t/myApp/myTrigger HTTP/1.1
Host: fnservice.com
Accept application/json 
```

Simple GET: (no body )
```
{
    "cloudEventsVersion" : "0.1",
    "eventType" : "io.fnproject.httpRequest",
    "eventTypeVersion" : "1.0",
    "source" : "https://fnservice.com/t/myApp/mytrigger",
    "eventID" : "A234-1234-1234",
    "eventTime" : "2018-04-05T17:31:00Z",
    "extensions" : {
        "ioFnProjectHttpReq" : {
            "method":"GET",
            "headers": {
                "Content-type" : ["application/json"],
             },
             "requestURL" : "https://fnservice.com/t/myApp/mytrigger"
         }
    }
}
```


## HTTP Response 

When a container returns with a cloud event to an HTTP gateway, the `data` of the event will be rendered as the HTTP body, and the `contentType` of the event will correspond to the returned HTTP content type.
 
By default all response events other than error events (see above) and annotated HTTP response events (see below) will be displayed with a `200` HTTP status code in the case where a `data` element is set or a `204` status code in the case here no `data` element is set.  



HTTP response metadata may be attached to any  returned event via the `ioFnProjectHttpResp` extension

 `"status": 200` : an HTTP status code to set on the response  
 `"headers" : {"header" : ["value1","value2"] }` : additional HTTP headers to apply to the response 

The body  o the HTTP response is derrived from the `data` element of the returned event, and the HTTP `Content-Type` of the HTTP response is derrived from the `contentType` of the returned event. Where a `Content-Type` header is specified in `ioFnProjectHttpResp.headers` extension, this is ignored by the HTTP gateway.   

The `ioFnProjectHttpResp` extension may be used on any returned event, however a specific event type `io.fnproject.httpResponse` is provided where no specific event type is impllied by the function 

e.g. Event returned by returned by container: 
```
{
    "cloudEventsVersion" : "0.1",
    "eventType" : "io.fnproject.httpResponse",
    "eventTypeVersion" : "1.0",
    "extensions" : {
        "ioFnProjectHttpResp" : {
            "status" : 204,
            "Content-type": "text/plain" // ignored 
        
        }
    }
    "contentType": "application/json; charset=utf-8",
    "data" : {"goto":"https://github.com/fnproject/fn","hello":"world!"}
   
}
```


```
HTTP/1.1 200 OK 
Content-Type: application/json; charset=utf-8
Date: Thu, 05 Jul 2018 16:36:47 GMT
Content-Length: 59

{"goto":"https://github.com/fnproject/fn","hello":"world!"}

```

