Http Servers:

        basic http server in go uses net/http:

        1. further net/http uses handlers  (an object implementing the http.Handler interface).

        2. common way to write ea handler is by using http.HandleFunc

        3. function serving as 

