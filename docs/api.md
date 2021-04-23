# API

Lo que es la api es bastante sencilla de manejar
pero primero debemos de saber un poco sobre las rutas que estara manejando el servidor ,
entre ellas esta:
## public
```
/public/{file:[\/\w\d\W]+?}
```

La cual es para acceder a los archivos que esten dentro de la carpeta `public` 

## publications
Despues tenemos otra ruta donde lo que vamos a ver son las publicaciones que se han hecho 
```
/{id:[0-9]+}
```
Y para ver una de las publicaciones debemos de ir a esta otra ruta
```
/publication/{id:[0-9]+}
```
Esta ruta funciona como un template donde se podran ver la publicacion con el id que se ha ingresado


## Post admin file
Esta permite el metodo `POST` y el metodo `GET` a la vez.
```
/admin/postfile
```
En el caso de que se quiera postear un nuevo post siempre y cuando tu ip encriptada con `sha256`este en la base de datos.
En estos casos el request debe de estar en un json.
Algo asi:
```json
{
  "title"            : "el titulo",
  "mineatura"        : "url de la imagen",
  "bodyOfDocument"   : "el markdown del archivo que se ha ingresado",
}
```
## finalmente la api

En lo que es la ruta de `api` solo hace falta hacer un request y este te regresara un json
```
/api/{page:[0-9]+}
```
El cuerpo del json es algo asi

```json
 {
  "Quantity": 9,
  "Publications": [
        {    
            "id": 1,
            "title": "hola mundo",
            "mineatura": "https://www.programaresfacil.co/wp-content/uploads/2018/02/Hola-Mundo.png",
            "bodyOfDocument": "# hello world"
        }
    ],
  "Size": 1
}
```
Es  bastante sencillo de manejar y lo que es cada parte es muy simple de explicar 

### Quantity

lo que es el atributo Quantity es para saber la cantidad de publicaciones que esten en cada homepage
 
### Size
Solo es para saber cuantas publicaciones hay disponibles

### Publications
Si quieres acceder a alguno de sus atributos como para ponerle alguna mineatura , titulo o manejar los links hacia las publicaciones pues para esto existe este atributo
