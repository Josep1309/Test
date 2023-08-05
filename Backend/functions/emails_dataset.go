package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"strings"
    "regexp"
)

type Email struct {
    Message_ID string `json:"Message-ID"`
    Date string `json:"Date"`
    From string `json:"From"`
    To string `json:"To"`
    Subject string `json:"Subject"`
    Cc string `json:"Cc"`
    Bcc string `json:"Bcc"`
    Mime_version string `json:"Mime-Version"`
    Content_type string `json:"Content-Type"`
    Content_transfer_encoding string `json:"Content-Transfer-Encoding"`
    X_from string `json:"X-From"`
    X_to string `json:"X-To"`
    X_cc string `json:"X-cc"`
    X_bcc string `json:"X-bcc"`
    X_folder string `json:"X-Folder"`
    X_origin string `json:"X-Origin"`
    X_filename string `json:"X-FileName"`
    Content string `json:"Content"`
}

func EmailsAdd(path string) {
    data, err := os.OpenFile("emails.ndjson", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        if os.IsNotExist(err) {
            data, err = os.Create("emails.ndjson")
            if err != nil {
                fmt.Println("Error creating file:", err)
                return
            }
        } else {
            fmt.Println("Error opening file:", err)
            return
        }
    }
	defer data.Close()

    encoder := json.NewEncoder(data)

    files, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if file.IsDir() == true {
            var folderPath string = path+file.Name()+"/"
            EmailsAdd(folderPath)
        }
        if file.IsDir() == false {
            var filePath string = path+file.Name()
            ReadFile(filePath, encoder, data)
        }
    }
}

func ReadFile(path string, encoder *json.Encoder, file *os.File){

    message := &Email{}

    readFile, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    re := regexp.MustCompile(`^(\w|\W)+:\s{2,}(\d|\D)+.*`)

    flag := -1

    for _, line := range fileLines {

        if strings.HasPrefix(line, "Message-ID:") && !re.MatchString(line) && flag<1{
            text, _:= strings.CutPrefix(line, "Message-ID:")
            text = strings.TrimLeft(text, " ")
            message.Message_ID = text
            flag = 0
        } else if strings.HasPrefix(line, "Date:") && !re.MatchString(line) && flag<2{
            text, _:= strings.CutPrefix(line, "Date:")
            text = strings.TrimLeft(text, " ")
            message.Date = text
            flag = 1
        } else if strings.HasPrefix(line, "From:" ) && !re.MatchString(line) && flag<3{
            text, _:= strings.CutPrefix(line, "From:")
            text = strings.TrimLeft(text, " ")
            message.From = text
            flag = 2
        } else if strings.HasPrefix(line, "To:") && !re.MatchString(line) && flag<4{
            text, _:= strings.CutPrefix(line, "To:")
            text = strings.TrimLeft(text, " ")
            message.To = text
            flag = 3
        } else if strings.HasPrefix(line, "Subject:") && !re.MatchString(line) && flag<5{
            text, _:= strings.CutPrefix(line, "Subject:")
            text = strings.TrimLeft(text, " ")
            message.Subject = text
            flag = 4
        } else if strings.HasPrefix(line, "Cc:") && !re.MatchString(line) && flag<6{
            text, _:= strings.CutPrefix(line, "Cc:")
            text = strings.TrimLeft(text, " ")
            message.Subject = text 
            flag = 5
        } else if strings.HasPrefix(line, "Bcc:") && !re.MatchString(line) && flag<7{
            text, _:= strings.CutPrefix(line, "Bcc:")
            text = strings.TrimLeft(text, " ")
            message.Subject = text
            flag = 6
        } else if strings.HasPrefix(line, "Mime-Version:") && !re.MatchString(line) && flag<8{
            text, _:= strings.CutPrefix(line, "Mime-Version:")
            text = strings.TrimLeft(text, " ")
            message.Mime_version = text
            flag = 7
        } else if strings.HasPrefix(line, "Content-Type:") && !re.MatchString(line) && flag<9{
            text, _:= strings.CutPrefix(line, "Content-Type:")
            text = strings.TrimLeft(text, " ")
            message.Content_type = text
            flag = 8
        } else if strings.HasPrefix(line, "Content-Transfer-Encoding:") && !re.MatchString(line) && flag<10{
            text, _:= strings.CutPrefix(line, "Content-Transfer-Encoding:")
            text = strings.TrimLeft(text, " ")
            message.Content_transfer_encoding = text
            flag = 9
        } else if strings.HasPrefix(line, "X-From:" ) && !re.MatchString(line) && flag<11{
            text, _:= strings.CutPrefix(line, "X-From:")
            text = strings.TrimLeft(text, " ")
            message.X_from = text
            flag = 10
        } else if strings.HasPrefix(line, "X-To:") && !re.MatchString(line) && flag<12{
            text, _:= strings.CutPrefix(line, "X-To:")
            text = strings.TrimLeft(text, " ")
            message.X_to = text
            flag = 11
        } else if strings.HasPrefix(line, "X-cc:") && !re.MatchString(line) && flag<13{
            text, _:= strings.CutPrefix(line, "X-cc:")
            text = strings.TrimLeft(text, " ")
            message.X_cc = text
            flag = 12
        } else if strings.HasPrefix(line, "X-bcc:") && !re.MatchString(line) && flag<14{
            text, _:= strings.CutPrefix(line, "X-bcc:")
            text = strings.TrimLeft(text, " ")
            message.X_bcc = text
            flag = 13
        } else if strings.HasPrefix(line, "X-Folder:") && !re.MatchString(line) && flag<15{
            text, _:= strings.CutPrefix(line, "X-Folder:")
            text = strings.TrimLeft(text, " ")
            message.X_folder = text
            flag = 14
        } else if strings.HasPrefix(line, "X-Origin:") && !re.MatchString(line) && flag<16{
            text, _:= strings.CutPrefix(line, "X-Origin:")
            text = strings.TrimLeft(text, " ")
            message.X_origin = text
            flag = 15
        } else if strings.HasPrefix(line, "X-FileName:") && !re.MatchString(line) && flag<17{
            text, _:= strings.CutPrefix(line, "X-FileName:")
            text = strings.TrimLeft(text, " ")
            message.X_filename = text
            flag = 16
        } else {
            switch flag {
            case 0:
                message.Message_ID = message.Message_ID + line
            case 1:
                message.Date = message.Date + line
            case 2:
                message.From = message.From + line
            case 3:
                message.To = message.To + line
            case 4:
                message.Subject = message.Subject + line
            case 5:
                message.Cc = message.Cc + line
            case 6:
                message.Bcc = message.Bcc + line
            case 7:
                message.Mime_version = message.Mime_version + line
            case 8:
                message.Content_type = message.Content_type + line
            case 9:
                message.Content_transfer_encoding = message.Content_transfer_encoding + line
            case 10:
                message.X_from = message.X_from + line
            case 11:
                message.X_to = message.X_to + line
            case 12:
                message.X_cc = message.X_cc + line
            case 13:
                message.X_bcc = message.X_bcc + line
            case 14:
                message.X_folder = message.X_folder + line
            case 15:
                message.X_origin = message.X_origin + line
            case 16:
                message.Content = message.Content + line + "\n"
            default:
                message.Content = message.Content + line + "\n"
            }
        }
    }

    if message.Message_ID != "" && message.Date != ""  {

        _, err := file.WriteString(`{ "index" : { "_index" : "emails" } }`)
        _, err = file.WriteString("\n")
		if err != nil {
			fmt.Println("Error encoding index:", err)
			return
		}

        emailErr := encoder.Encode(message)
		if emailErr != nil {
			fmt.Println("Error encoding email:", err)
			return
		}
    }
}