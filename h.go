package main

import (
    "github.com/integrii/flaggy"
    "log"
	"net/http"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    var headers bool
    var body bool
    var status bool
    var tls bool
    var chrome bool
    var ie6 bool
    var safari bool
    var edge bool
    var malware bool
    var bogusff bool
    var openvas bool
    var meterpreter bool
    var save =""
    var proxy = ""
    var uagent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0"
    var url = "http://www.google.com"
    flaggy.Bool(&headers,"x","headers" ,"print the http headers of the http response")
    flaggy.Bool(&body,"b", "body","print the body of the http response")
    flaggy.Bool(&status,"s", "status","print the status code of the http response")
    flaggy.Bool(&tls,"t", "tls","choose if the request is over tls")
    flaggy.Bool(&chrome,"c", "chrome" ,"Use Chrome on Mac User Agent")
    flaggy.Bool(&ie6,"i", "ie6" ,"Use Internet Explorer 6.0 User Agent")
    flaggy.Bool(&safari,"m", "safari" ,"Use Mac Safari User Agent")
    flaggy.Bool(&edge,"e", "edge" ,"Use Edge User Agent")
    flaggy.Bool(&malware,"v", "malware" ,"Use known bad User Agent malware")
    flaggy.Bool(&bogusff,"y", "bogusff","Use known bad User Agent Firefox/1")
    flaggy.Bool(&openvas,"o", "openvas","Use known bad User Agent OpenVAS")
    flaggy.Bool(&meterpreter,"z", "meterpreter","Use known hacking tool User Agent Meterpreter")
    flaggy.String(&save,"f","save" ,"Save the body to a file")
    flaggy.String(&proxy,"p","proxy" ,"Specify the IP address and port of proxy - 127.0.0.1:8080")
    flaggy.String(&uagent,"a", "uagent" ,"Specify a custom user-agent in quotes")
    flaggy.String(&url,"u", "url" ,"Enter a url with or without leading http:// or https://")
    flaggy.Parse()
    
    if headers!=true && body!=true && status!=true {
        body = true   
    }
    
    if strings.Contains(url, ":"){
        s := strings.Split(url, ":")
        if s[0] == "http" && tls {
            url=("https:" + strings.Join(s[1:]," "))
        }
    } else {
        if tls{
            url=("https://" + url)
        }else {
            url=("http://" + url)
        }

    }
    
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "close")
    
    if ie6 {
    uagent="Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727)"
    } else if chrome {
    uagent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
    } else if safari {
    uagent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.8 (KHTML, like Gecko)"
    } else if edge{
    uagent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063"
    } else if meterpreter {
        uagent="Meterpreter"
    } else if malware {
        uagent="malware"
    } else if openvas {
        uagent="OpenVAS"
    } else if bogusff { 
        uagent=" Firefox/1."
    } 
    if len(uagent) > 0 {
        req.Header.Set("User-Agent", uagent) }
    
    print("\nURL: ", url,"\n")
    print("\nUSER-AGENT: ",uagent,"\n")
    
    if len(proxy) > 1 {
        os.Setenv("HTTP_PROXY", proxy)
    }
    
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
    responseData, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    responseString := string(responseData)
    
    if status {
        print ("\nRESPONSE CODE: ", (res.StatusCode), "\n") 
    }
    if headers {
        print ("\n========= HEADERS ==========\n")
        for k, v := range res.Header {
            fmt.Print(k, v,"\n")
        }
    }
    if len(save) > 0 {
        s, err := os.Create(save)
        if err != nil {
            fmt.Println(err)
            return
    }
        b, err := s.WriteString(responseString)
        if err != nil {
            fmt.Println(err)
            s.Close()
            return
    }
        fmt.Println("\n============ File Saved =============")
        fmt.Println(b, "bytes were successfully written to ",save,"\n\n")
        err = s.Close()
        if err != nil {
            fmt.Println(err)
            return
    }
    } else if body {
        fmt.Println("\n============ BODY =============\n", responseString)
    }
    
    
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
}
    print("\n")
}