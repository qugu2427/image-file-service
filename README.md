# GoImageServer
An image storage server written in Go.

Can be used as a general file server, but is specifcally tailored towards image storage.
Designed to run in a local network and be used by a backend.

To Use: <br>
-Set desired upload folder in config.go <br>
-Build project or "go run main.go config.go" in terminal <br>

POST /upload <br>
Formdata: "file" <br>
Uploads file.

DELETE /delete <br>
Formdata: "fileName" <br>
Deletes specified file.
Will always return 200 even is the file doesn't exist.

ALL /<filename> <br>
Gets file.
