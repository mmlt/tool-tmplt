package thycotic

import (
	"testing"
	"encoding/xml"
	"fmt"
)

//
func Unmarchall_GetSecretResponse(t *testing.T) {
	rawbody := []byte(`
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
               xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <soap:Body>
        <GetSecretResponse xmlns="urn:thesecretserver.com">
            <GetSecretResult>
                <Errors/>
                <Secret>
                    <Name>artifactory.otas.nv replication user</Name>
                    <Items>
                        <SecretItem>
                            <Value/>
                            <Id>170102</Id>
                            <FieldId>60</FieldId>
                            <FieldName>Server</FieldName>
                            <IsFile>false</IsFile>
                            <IsNotes>false</IsNotes>
                            <IsPassword>false</IsPassword>
                            <FieldDisplayName>Resources</FieldDisplayName>
                        </SecretItem>
                        <SecretItem>
                            <Value>replication</Value>
                            <Id>170103</Id>
                            <FieldId>61</FieldId>
                            <FieldName>Username</FieldName>
                            <IsFile>false</IsFile>
                            <IsNotes>false</IsNotes>
                            <IsPassword>false</IsPassword>
                            <FieldDisplayName>Username</FieldDisplayName>
                        </SecretItem>
                        <SecretItem>
                            <Value>this-is-a-secret</Value>
                            <Id>170104</Id>
                            <FieldId>7</FieldId>
                            <FieldName>Password</FieldName>
                            <IsFile>false</IsFile>
                            <IsNotes>false</IsNotes>
                            <IsPassword>true</IsPassword>
                            <FieldDisplayName>Password</FieldDisplayName>
                        </SecretItem>
                        <SecretItem>
                            <Value>artifactory.otas.nv replication user</Value>
                            <Id>170105</Id>
                            <FieldId>8</FieldId>
                            <FieldName>Notes</FieldName>
                            <IsFile>false</IsFile>
                            <IsNotes>true</IsNotes>
                            <IsPassword>false</IsPassword>
                            <FieldDisplayName>Notes</FieldDisplayName>
                        </SecretItem>
                    </Items>
                    <Id>37027</Id>
                    <SecretTypeId>2</SecretTypeId>
                    <FolderId>496</FolderId>
                    <IsWebLauncher>false</IsWebLauncher>
                    <CheckOutMinutesRemaining xsi:nil="true"/>
                    <IsCheckedOut xsi:nil="true"/>
                    <CheckOutUserDisplayName/>
                    <CheckOutUserId xsi:nil="true"/>
                    <IsOutOfSync xsi:nil="true"/>
                    <IsRestricted>false</IsRestricted>
                    <OutOfSyncReason/>
                    <Active>true</Active>
                </Secret>
            </GetSecretResult>
        </GetSecretResponse>
    </soap:Body>
</soap:Envelope>
	`)
	response := new(GetSecretResponse)
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err := xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		fmt.Print(err)
	}
}
