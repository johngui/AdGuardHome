package home

import (
	"testing"
	"time"
)

/* Tests performed:
. Bad certificate
. Bad private key
. Valid certificate & private key */
func TestValidateCertificates(t *testing.T) {
	var data tlsConfigStatus

	// bad cert
	data = validateCertificates("bad cert", "", "")
	if !(data.WarningValidation != "" &&
		!data.ValidCert &&
		!data.ValidChain) {
		t.Fatalf("bad cert: validateCertificates(): %v", data)
	}

	// bad priv key
	data = validateCertificates("", "bad priv key", "")
	if !(data.WarningValidation != "" &&
		!data.ValidKey) {
		t.Fatalf("bad priv key: validateCertificates(): %v", data)
	}

	// valid cert & priv key
	CertificateChain := `-----BEGIN CERTIFICATE-----
MIICKzCCAZSgAwIBAgIJAMT9kPVJdM7LMA0GCSqGSIb3DQEBCwUAMC0xFDASBgNV
BAoMC0FkR3VhcmQgTHRkMRUwEwYDVQQDDAxBZEd1YXJkIEhvbWUwHhcNMTkwMjI3
MDkyNDIzWhcNNDYwNzE0MDkyNDIzWjAtMRQwEgYDVQQKDAtBZEd1YXJkIEx0ZDEV
MBMGA1UEAwwMQWRHdWFyZCBIb21lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB
gQCwvwUnPJiOvLcOaWmGu6Y68ksFr13nrXBcsDlhxlXy8PaohVi3XxEmt2OrVjKW
QFw/bdV4fZ9tdWFAVRRkgeGbIZzP7YBD1Ore/O5SQ+DbCCEafvjJCcXQIrTeKFE6
i9G3aSMHs0Pwq2LgV8U5mYotLrvyFiE8QPInJbDDMpaFYwIDAQABo1MwUTAdBgNV
HQ4EFgQUdLUmQpEqrhn4eKO029jYd2AAZEQwHwYDVR0jBBgwFoAUdLUmQpEqrhn4
eKO029jYd2AAZEQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQB8
LwlXfbakf7qkVTlCNXgoY7RaJ8rJdPgOZPoCTVToEhT6u/cb1c2qp8QB0dNExDna
b0Z+dnODTZqQOJo6z/wIXlcUrnR4cQVvytXt8lFn+26l6Y6EMI26twC/xWr+1swq
Muj4FeWHVDerquH4yMr1jsYLD3ci+kc5sbIX6TfVxQ==
-----END CERTIFICATE-----`
	PrivateKey := `-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALC/BSc8mI68tw5p
aYa7pjrySwWvXeetcFywOWHGVfLw9qiFWLdfESa3Y6tWMpZAXD9t1Xh9n211YUBV
FGSB4ZshnM/tgEPU6t787lJD4NsIIRp++MkJxdAitN4oUTqL0bdpIwezQ/CrYuBX
xTmZii0uu/IWITxA8iclsMMyloVjAgMBAAECgYEAmjzoG1h27UDkIlB9BVWl95TP
QVPLB81D267xNFDnWk1Lgr5zL/pnNjkdYjyjgpkBp1yKyE4gHV4skv5sAFWTcOCU
QCgfPfUn/rDFcxVzAdJVWAa/CpJNaZgjTPR8NTGU+Ztod+wfBESNCP5tbnuw0GbL
MuwdLQJGbzeJYpsNysECQQDfFHYoRNfgxHwMbX24GCoNZIgk12uDmGTA9CS5E+72
9t3V1y4CfXxSkfhqNbd5RWrUBRLEw9BKofBS7L9NMDKDAkEAytQoIueE1vqEAaRg
a3A1YDUekKesU5wKfKfKlXvNgB7Hwh4HuvoQS9RCvVhf/60Dvq8KSu6hSjkFRquj
FQ5roQJBAMwKwyiCD5MfJPeZDmzcbVpiocRQ5Z4wPbffl9dRTDnIA5AciZDthlFg
An/jMjZSMCxNl6UyFcqt5Et1EGVhuFECQQCZLXxaT+qcyHjlHJTMzuMgkz1QFbEp
O5EX70gpeGQMPDK0QSWpaazg956njJSDbNCFM4BccrdQbJu1cW4qOsfBAkAMgZuG
O88slmgTRHX4JGFmy3rrLiHNI2BbJSuJ++Yllz8beVzh6NfvuY+HKRCmPqoBPATU
kXS9jgARhhiWXJrk
-----END PRIVATE KEY-----`
	data = validateCertificates(CertificateChain, PrivateKey, "")
	notBefore, _ := time.Parse(time.RFC3339, "2019-02-27T09:24:23Z")
	notAfter, _ := time.Parse(time.RFC3339, "2046-07-14T09:24:23Z")
	if !(data.WarningValidation != "" /* self signed */ &&
		data.ValidCert &&
		!data.ValidChain &&
		data.ValidKey &&
		data.KeyType == "RSA" &&
		data.Subject == "CN=AdGuard Home,O=AdGuard Ltd" &&
		data.Issuer == "CN=AdGuard Home,O=AdGuard Ltd" &&
		data.NotBefore == notBefore &&
		data.NotAfter == notAfter &&
		// data.DNSNames[0] ==  &&
		data.ValidPair) {
		t.Fatalf("valid cert & priv key: validateCertificates(): %v", data)
	}
}
