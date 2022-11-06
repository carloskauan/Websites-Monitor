package main

import (
  "fmt"
  "os"
  "bufio"
  "net/http"
  "io"
  "strings"
  "io/ioutil"
  "time"
  "strconv"
)

func main(){
  os.Create("logs.txt")
  var choice int
  for{
    var count int
    fmt.Println("Monitorador de sites\n\n1-\tMonitorar\n2-\tRegistrar sites\n3-\tExibir logs\n\n0-\tSair")
    fmt.Scan(&choice)
    switch choice{
    case 1:
      txtSites, _ := os.OpenFile("sites.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
      readSiteTxt := bufio.NewReader(txtSites)
      for{
        readSite, errR := readSiteTxt.ReadString('\n')
        if errR == io.EOF{
          if count != 0{
            break
          }else{
            fmt.Println("\nNem um site registrado...\nUse a opção Registrar sites no menu <2>\n")
            time.Sleep(3*time.Second)
            break
          }
        }
        readSite = strings.TrimSpace(readSite)
        resp, _ := http.Get(readSite)
        switch resp.StatusCode{
        case 200:
          registerLog(readSite, true)
          fmt.Println("Status: Carregado com sucesso\nUrl:",readSite,"\nStatus code:",resp.StatusCode,"\n")
        default:
          registerLog(readSite, false)
          fmt.Println("Status: Falha ao carregar\nUrl:",readSite,"\nStatus code:",resp.StatusCode,"\n")
        }
        count = 1
      }
      txtSites.Close()
    case 2:
      var quant int
      var site string
      txtSites, _ := os.OpenFile("sites.txt", os.O_WRONLY|os.O_APPEND, 0666)
      fmt.Println("\nRegistro de sites\n\nQuantos sites deseja registrar?")
      fmt.Scan(&quant)
      for c := 1; c <= quant; c++{
        fmt.Println("Registro -", c)
        fmt.Scan(&site)
        txtSites.WriteString(site+"\n")
      }
      fmt.Println("Sites Registrados com sucesso")
    case 3:
      txtLogs, _ := ioutil.ReadFile("logs.txt")
      if txtLogs == nil {
        fmt.Println("Sem logs registrados")
      }else{
        fmt.Println(string(txtLogs))
      }
    case 0:
      os.Exit(0)
    }
  }
}
func registerLog(url string, status bool){
  txtLogs, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
  if err == io.EOF{
    fmt.Println("Nem um log registrado")
  }else{
    txtLogs.WriteString(url+" Online: "+strconv.FormatBool(status)+" Horario: "+time.Now().Format("02/01/2006 15:04:05")+"\n")
  }
}
