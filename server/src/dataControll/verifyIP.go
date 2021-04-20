package dataControll

import "github.com/ranon-rat/silent-songs/src/stuff"

func VerifyCookieIP(ip string,canPublic chan bool){
	q:=`
	SELECT COUNT(*) FROM  users WHERE ip_encrypted=?1;
	`
	db:=GetConnection()
	r,_:=db.Query(q,stuff.EncryptData(ip))
	howMany:=0
	for r.Next(){
		r.Scan(&howMany)
	}
	canPublic<-howMany>0

}