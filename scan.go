//this programme ssh into the provided list of ip addresses and saves the result of mongo query result to /tmp/supportip.txt
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	type MyStruct struct {
		Name string `json:"name"`
	}

	var b bytes.Buffer
	var c bytes.Buffer
	var result map[string]interface{}
	var Tmpstr MyStruct
	var Orgname string
	//Tmpsli := []string{}
        keypath := os.Args[1]
	hostiplist := []string {"20.19.1.241","20.19.0.71","20.19.0.72","20.19.0.73"}
	port := "22"
	key,err := ioutil.ReadFile(keypath)
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}
	signer,err:=ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v",err)
	}
    var conn *ssh.Client

	config := &ssh.ClientConfig{
		User:            "soldier",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 5*time.Second,
	}
	//loop starts here
    for _,hostip := range hostiplist {
		b.Reset()  //reset the buffer other wise next query result will be appended here
		c.Reset()
		fmt.Printf("%T",c)
		conn,err = ssh.Dial("tcp",net.JoinHostPort(hostip,port),config)
		if err != nil {
			log.Fatalf("unable to connect: %v",err)
		}
		defer conn.Close()
		session1,err:=conn.NewSession()
		if err != nil {
			log.Fatalf("unable to connect: %v",err)
		}
		defer session1.Close()
		session1.Stdout = &b
        //escape special character in password e.g.. -u mongouser -p \"s3creTwo34Five\"

		if err := session1.Run("/usr/bin/test -f /usr/bin/mongo && /usr/bin/mongo mydbname -u myuser -p xxxxxx --quiet --eval \"db.customer.findOne({},{\"_id\":0,\"name\":1})\" || /opt/turtle/bin/mongo mydbname -u mydbname -p xxxxxxx --quiet --eval \"db.customer.findOne({},{\"_id\":0,\"name\":1})\""); err != nil {
			log.Fatal("Failed to run:sess1 " + err.Error())
		}
		fmt.Printf("%v %v \n",hostip,b.String())
		err1 := json.Unmarshal([]byte(b.String()),&Tmpstr)
		if err1 != nil {
			log.Fatalf("Failed Unmarshal1 %v",err1)
		}
		fmt.Println(Tmpstr.Name)

		session2,err:=conn.NewSession()
		session2.Stdout = &c
		if err := session2.Run("/usr/bin/test -f /usr/bin/mongo && /usr/bin/mongo mydbname -u myuser -p xxxxxx  --quiet --eval \"db.license.findOne({},{ \"_id\" : 0, \"customer\" : 1 })\" || /opt/turtle/bin/mongo mydbname -u mydbname -p xxxxxxx  --quiet --eval \"db.license.findOne({},{ \"_id\" : 0, \"customer\" : 1 })\" "); err != nil {
			log.Fatal("Failed to run:sess2 " + err.Error())
		}
		err2 := json.Unmarshal([]byte(c.String()),&result)
		if err2 != nil {
			log.Fatalf("Failed Unmarshal for Customer name %v",err2)
		}
		fmt.Println(c.String())
		custmap := result["customer"].(map[string]interface{})
		

		for key, value := range custmap {
			// Each value is an interface{} type, that is type asserted as a string
			//fmt.Println(key, value.(string))
			if key == "name" {
				Orgname=value.(string)
			}
		}
		fmt.Println("second query result: ",hostip,ServerName,Orgname)
		f, err := os.OpenFile("/tmp/supportip.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(hostip+","+Tmpstr.Name+","+Orgname+"\n")); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

}
