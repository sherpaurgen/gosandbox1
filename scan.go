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

type hostInfo struct {
	Name string
}

var b bytes.Buffer
var c bytes.Buffer

func main() {

	hostip := "192.168.1.150"

	port := "22"

	key,err := ioutil.ReadFile("/path/to/pvtkey")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer,err:=ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v",err)
	}

	config := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 5*time.Second,
	}
	conn,err := ssh.Dial("tcp",net.JoinHostPort(hostip,port),config)
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
	if err := session1.Run("some onliner mongo query.."); err != nil {
		log.Fatal("Failed to run session1 cmd" + err.Error())
	}
	var tmphostInfo hostInfo
    hostjsonstr:=b.String()
	err1 := json.Unmarshal([]byte(hostjsonstr),&tmphostInfo)
	if err1 != nil {
		log.Fatalf("Failed Unmarshal1 %v",err1)
	}


    session2,err:=conn.NewSession()
	session2.Stdout = &c
	if err := session2.Run("another oneliner db query"); err != nil {
		log.Fatal("Failed to run session2 cmd" + err.Error())
	}
	var result map[string]interface{}
	infojson:=c.String()
	json.Unmarshal([]byte(infojson),&result)
    //fmt.Println(c.String())
	birds := result["customer"].(map[string]interface{})
	//fmt.Println(birds)
    var LPname string
	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		//fmt.Println(key, value.(string))
		if key == "name" {
			LPname=value.(string)
		}
	}
	fmt.Println(hostip,tmphostInfo.Name,LPname)
	f, err := os.OpenFile("/tmp/vpnip.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(hostip+","+tmphostInfo.Name+","+LPname+"\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
