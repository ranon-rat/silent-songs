# API

What is the api is quite simple to handle
But first we must know a little about the routes that the server will be handling,
among them is:
## public
```
/ public / {file: [\ / \ w \ d \ W] +?}
```

Which is to access the files that are inside the`public` folder

## publications
Then we have another route where what we are going to see are the publications that have been made
```
/ {id: [0-9] +}
```
And to see one of the publications we must go to this other route
```
/ publication / {id: [0-9] +}
```
This route works as a template where you can see the publication with the id that has been entered


## Post admin file
It allows the`POST` method and the` GET` method at the same time.
```
/ admin / postfile
```
In case you want to post a new post as long as your ip encrypted with`sha256` is in the database.
In these cases the request must be in a json.
Something like that:
```json
{
  "title": "the title",
  "miniature": "image url",
  "bodyOfDocument": "the markdown of the file that has been entered",
}
```
## finally the api

In what is the`api` route, you only need to make a request and it will return a json
```
/ api / {page: [0-9] +}
```
The body of the json is something like this

```json
 {
  "Quantity": 9,
  "Publications": [
        {
            "id": 1,
            "title": "hello world",
            "miniature": "https://www.programaresfacil.co/wp-content/uploads/2018/02/Hola-Mundo.png",
            "bodyOfDocument": "# hello world"
        }
    ],
  "Size": 1
}
```
It is quite simple to handle and what each part is is very simple to explain

### Quantity

what is the Quantity attribute is to know the number of publications that are on each homepage
 
### Size
It is only to know how many publications are available

### Publications
If you want to access any of its attributes such as to put a miniature, title or manage the links to the publications, for this there is this attribute