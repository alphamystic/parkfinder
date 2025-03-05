package main

import (
 "fmt"
 "net/http"
 "crypto/tls"
 "ken/lib/utils"
 "ken/lib/ui/router"
 //ent"ken/ibgo/entities"
)

func main(){
 utils.PTSB("white","PARK VIEWER running at port 4000 (HTTP) and 4001 (HTTPS).")
 pfl := &SoilSampling {
   Address: "0.0.0.0",
   PortS: 4001,
   Port: 4000,
   // TlsCert string
   // TlsKey string
   Tls: false,
 }
 var svr = new(http.Server)
 var svrs = new(http.Server)
 svr, err := pfl.CreateServer()
 if err != nil {
   panic(err)
 }
 pfl.Tls = true
 svrs,err = pfl.CreateServer()
 if err != nil {
   panic(err)
 }
 //mux := http.NewServeMux()
 rtr := router.NewRouter(svrs,svr)
 rtr.Run(true)
}

type SoilSampling struct {
 Address string
 PortS int
 Port int
 TlsCert string
 TlsKey string
 Tls bool
}

func (l *SoilSampling) CreateServer() (*http.Server,error) {
 if l.Tls {
   config := &tls.Config {
     MinVersion: tls.VersionTLS12,
     CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
     PreferServerCipherSuites: true,
     CipherSuites: []uint16 {
       tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
       tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
       tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
       tls.TLS_RSA_WITH_AES_256_CBC_SHA,
     },
   }
   return &http.Server {
     Addr: fmt.Sprintf(":%d",l.PortS),
     TLSConfig: config,
     TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
   },nil
 } else {
   return &http.Server {
     Addr: fmt.Sprintf(":%d", l.Port),
   },nil
 }
 return nil,fmt.Errorf("You probably have an error in your SoilSampling initialization.")
}

// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
