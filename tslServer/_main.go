package main

import (
	"fmt"
	"crypto/x509"
	"math/big"
	"crypto/rand"
	"errors"
	"crypto/x509/pkix"
	"time"
	"crypto/rsa"
	"log"
	"net"
	"encoding/pem"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Generating random key: %v", err)
	}
	
	rootCertTmpl, err := CertTemplate()
	if err != nil {
		log.Fatalf("Creating cert template: %v", err)
	}
	
	rootCertTmpl.IsCA = true
	rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature
	rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
	rootCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
	
	rootCert, rootCertPEM, err := CreateCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("errpr creating cert: %v", err)
	}
	
	fmt.Printf("%s\n", rootCertPEM)
	fmt.Printf("%#x\n", rootCert.Signature)
	
	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})
	
	_, err = tls.X509KeyPair(rootCertPEM, rootKeyPEM)
	if err != nil {
		log.Fatalf("Invalid key pair: %v", err)
	}
	
	servKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Generate random key: %v", err)
	}
	
	servCertTmpl, err := CertTemplate()
	if err != nil {
		log.Fatalf("Creating cert template: %v", err)
	}
	
	servCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	servCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	servCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
	
	_, servCertPEM, err := CreateCert(servCertTmpl, rootCert, &servKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("Error creating cert: %v", err)
	}
	
	servKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(servKey),
	})
	
	servTLSCert, err := tls.X509KeyPair(servCertPEM, servKeyPEM)
	if err != nil {
		log.Fatalf("Invalid key pair: %v", err)
	}
	
	httpHandleFunc := func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("HI !"))
	}
	
	httpServer := httptest.NewUnstartedServer(http.HandlerFunc(httpHandleFunc))
	
	httpServer.TLS = &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
	}
	
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(rootCertPEM)
	
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
	}
	
	httpServer.StartTLS()
	res, err := client.Get(httpServer.URL)
	httpServer.Close()
	
	if err != nil {
		log.Fatalf("Could not make GET request: %v", err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatalf("Could not dump response: %s", err)
	}
	fmt.Printf("%s\n", dump)
}

func CertTemplate() (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("failed to generate serial number: "+ err.Error())
	}
	
	tmpl := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Danh Tran"},
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(time.Hour),
		BasicConstraintsValid: true,
	}
	return &tmpl, nil
}

func CreateCert(template, parent *x509.Certificate, pub interface{}, parentPriv interface{}) (cert *x509.Certificate, certPEM []byte, err error) {
	certDER, err := x509.CreateCertificate(rand.Reader, template, parent, pub, parentPriv)
	if err != nil {
		return
	}
	
	cert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return
	}
	
	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM = pem.EncodeToMemory(&b)
	return
}


