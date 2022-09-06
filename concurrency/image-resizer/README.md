# Image resizer

## Task

The task is to add a new URL parameter `async`. If the new parameter is set to `true` we process images in background and return the response imidiatelly.

The `getImageHandler` function should wait for finishing the processing the image if it's not yet available.

## Build & Run Server locally
```
go run .
```

## Run a sample request against the server
```
curl -X POST -H "Content-Type: application/json" -d @req.json http://localhost:8080/v1/resize
```

Now in your browser, you can check one of the returned urls!
