package main

import (
	"PorAddress/email"
	"PorAddress/ipaddr"
	"fmt"
)

func main() {
	al := &ipaddr.AddressList{}
	err := al.GetAddressList()
	if err != nil {
		fmt.Println("get address error")
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

	err = el.SendEmail(e_content)
	if err != nil {
		fmt.Println("send email error")
	}

}
