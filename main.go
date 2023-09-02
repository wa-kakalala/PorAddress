package main

import (
	"PorAddress/email"
	"PorAddress/ipaddr"
	"fmt"
	"time"
)

func main() {
	al := &ipaddr.AddressList{}
	for {
		err := al.GetAddressList()
		if err != nil {
			fmt.Println("get address error")
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("get address successfully")
			break
		}
	}

	el := &email.EmailInfo{
		To:      "xxxxxx@qq.com", // string
		From:    "xxxxxx@qq.com", // string
		Subject: "Device Info",   // string
		Host:    "smtp.qq.com",   // string
		Port:    465,             // int
		User:    "xxxxxx@qq.com", // string
		Passwd:  "xxxxxx",        // string
	}

	// emoji ： https://emojixd.com/group/smileys-emotion
	e_content := ""
	e_content += "Device is Starting...           \n"
	e_content += "😁==AddressInfo==\n"
	e_content += "--------------------------------\n"
	for index := 0; index < len(al.Addresses); index++ {
		e_content += al.NetNames[index]
		e_content += ":&nbsp;"
		e_content += al.Addresses[index]
		e_content += "\n"
	}
	e_content += "--------------------------------\n"
	e_content += "The End!!!👻"

	err := el.SendEmail(e_content)
	for {
		if err != nil {
			fmt.Println("send email error")
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("send email successfully")
			break
		}
	}

}
