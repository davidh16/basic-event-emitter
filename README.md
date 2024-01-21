# Description

This is simple event emitter written in GO. 
The purpose of this app is to be used for quick and easy testing of socket.io applications. 

It is based on github.com/graarh/golang-socketio package.

App is tested for the socket.io server which is also written in GO and is based on "github.com/googollee/go-socket.io" package.

Link to the socket.io server repository : xxx

## Usage

The app runs on HTTP protocol 
therefore it is necessary to send POST request to 
proper url to hit EmitEvent route which expects JSON that has the name of method (event that wants to be emitted) and a message.

Example :
```json
{
    "method": "my_event",
    "message": "hello there !"
}
```

It is possible to extend this HTTP server with more custom routes by writing custom route and registering it with the server router in NewServer() function.

Example :

 custom route logic
```go
func (s *Server) MyCustomRoute(w http.ResponseWriter, req *http.Request) {
	...
	my custom logic
	...
    
	
	err = s.client.Emit(method, message)
	if err != nil {
		s.returnResponse(w, http.StatusInternalServerError, err)
	}

	s.returnResponse(w, http.StatusOK, nil)
}
```

registering route with the server router
```go
func NewServer(cfg *config.Config, socketClient *gosocketio.Client) *Server {

	...
	server definition
	...
	
	router.HandleFunc("my_custom_route_url", srv.MyCustomRoute).Methods("WHATEVER_METHOD_ONE_WANTS") #POST/GET

	return srv

}
```

Application can be run inside Docker by creating docker-compose.yml file (use docker-compose-example.yml for the reference) and running `docker compose up` command.