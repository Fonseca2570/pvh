# pvh
pvh challenge


# compare service
Service runs on port 8800 and as the following endpoints

    - GET http://localhost:8800/alive -> checks if the service is running
    - GET http://localhost:8800/files/health -> checks if the service has both current.json and backup.json in the root file (default there is)
    - GET http://localhost:8800/files/:type -> type can either be current or backup other inputs will get error, just prints the content of the file to the client
    - PUT http://localhost:8800/upload/file/:type -> type can either be current or backup other inputs will get error, form data should be a json file that will overwrite the existing one defined in type key for the form should be "file"
    - GET http://localhost:8800/compare -> will compare the files in the system to check if there are equals if not it will give the diference between the files.


The repository has a postman_collection to import so that is easier to interact with the application

To run the app just use docker-compose up on the pvh folder

By default the backup and current files are equal but order is not the same