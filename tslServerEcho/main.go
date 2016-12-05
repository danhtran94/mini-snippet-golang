package main

import (
	"crypto/rsa"
	"crypto/rand"
	"log"
	"crypto/x509"
	"math/big"
	"github.com/kataras/go-errors"
	"crypto/x509/pkix"
	"time"
	"net"
	"encoding/pem"
	"os"
	"bytes"
	"fmt"
	"crypto/tls"
	"bufio"
	
	"net/http/httptest"
	"net/http/httputil"
	"net/http"
)

func main() {
	// Tạo một cặp key rsa có size 2048
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
	rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageClientAuth,
	}
	rootCertTmpl.IPAddresses = []net.IP{
		net.ParseIP("127.0.0.1"),
	}
	
	rootCert, rootCertPEM, err := CreateCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("Error creating cert: %v", err)
	}
	
	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})
	
	err = SavePEM("rootCert.pem", rootCertPEM)
	if err != nil {
		log.Fatalf("Error saving pem: %v", err)
	}
	
	err = SavePEM("rootKey.pem", rootKeyPEM)
	if err != nil {
		log.Fatalf("Error saving pem: %v", err)
	}
	
	rootCertPEM, err = OpenPEM("rootCert.pem")
	rootKeyPEM, err = OpenPEM("rootKey.pem")
	if err != nil {
		log.Fatalf("Error opening pem: %v", err)
	}
	
	_, err = tls.X509KeyPair(rootCertPEM, rootKeyPEM)
	if err != nil {
		log.Fatalf("Invalid key pair: %v", err)
	}
	
	// Server
	servKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Generate random key: %v", err)
	}
	
	servKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(servKey),
	})
	
	err = SavePEM("servKey.pem", servKeyPEM)
	if err != nil {
		log.Fatalf("Error saving pemfile: %v", err)
	}
	
	servCertTmpl, err := CertTemplate()
	if err != nil {
		log.Fatalf("Creating cert template: %v", err)
	}
	
	servCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
	servCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
	}
	servCertTmpl.IPAddresses = []net.IP{
		net.ParseIP("127.0.0.1"),
	}
	
	_, servCertPEM, err := CreateCert(servCertTmpl, rootCert, &servKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("Error creating server cert: %v", err)
	}
	
	err = SavePEM("servCert.pem", servCertPEM)
	if err != nil {
		log.Fatalf("Error saving server cert: %v", err)
	}
	
	servCertPEM, err = OpenPEM("servCert.pem");
	servKeyPEM, err = OpenPEM("servKey.pem");
	if err != nil {
		log.Fatalf("Invalid key pair: %v", err)
	}
	
	servTLSCert, err := tls.X509KeyPair(servCertPEM, servKeyPEM)
	if err != nil {
		log.Fatalf("Invalid key pair: %v", err)
	}
	
	///* http server
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
	//*/
}

func OpenPEM(fileName string) ([]byte, error) {
	filePem, err := os.Open(fileName)
	var pemBytes [2000]byte
	n, _ := bufio.NewReader(filePem).Read(pemBytes[0:])
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	w, _ := buffer.Write(pemBytes[0:n])
	fmt.Println(w)
	return buffer.Bytes(), nil
}

func SavePEM(fileName string, pemBytes []byte) error {
	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	w, _ := buffer.Write(pemBytes)
	fmt.Println(w)
	
	filePem, err := os.Create(fileName)
	if err != nil {
		return err
	}
	
	wt, _ := buffer.WriteTo(filePem)
	fmt.Println(wt)
	return nil
}

func CertTemplate() (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("Failed to generate serial number: "+ err.Error())
	}
	
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Tran Thanh Danh Corp"},
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(time.Hour),
		BasicConstraintsValid: true,
	}
	
	return &template, nil
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
	
	block := pem.Block{
		Type: "CERTIFICATE",
		Bytes: certDER,
	}
	
	certPEM = pem.EncodeToMemory(&block)
	return
}