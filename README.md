# Forum

This project consists in creating a web forum that allows :

    communication between users.
    associating categories to posts.
    liking and disliking posts and comments.
    filtering posts.

### How to go run?

First, u need to enter backend directory

```
go run main.go

```

Next, enter the client directory and run these commands:

```
npm install
```

```
npm start
```

After that, React will run the application at localhost:3000 port

### How to use

Authorization is required prior to the usage of the forum itself. After you log in, you can create posts, comment other's posts, like and dislike them and etc.

### If you want to use makefile and docker, use this command###

```
make dcrun
```

Wait for the completion of the image creation

```
make dbrun
```

Wait for completion of the db image creation

# go to link localhost:3000/

if you want to delete all images, use this command

```
make dstop
```

```
make dclear
```

### Project made by

@aromanov team-lead

@diyar.ildart backend

@aseitkhan frontend
