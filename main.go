package main

import (
	"fmt"
	"log"
	"net/http"
)
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World Phani!")
}
func mockacct(w http.ResponseWriter, r *http.Request) {
	j := `{
    "d": {
        "__metadata": {
            "id": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')",
            "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')",
            "type": "ERP_FICA_UMC.Account"
        },
        "AccountTypeID": "1",
        "AccountID": "2000018",
        "AccountTitleID": "",
        "FirstName": "Gabrielle",
        "LastName": "Lynch",
        "MiddleName": "",
        "SecondName": "",
        "Sex": "",
        "Name1": "",
        "Name2": "",
        "Name3": "",
        "Name4": "",
        "GroupName1": "",
        "GroupName2": "",
        "FullName": "Gabrielle Lynch",
        "CorrespondenceLanguage": "E",
        "CorrespondenceLanguageISO": "EN",
        "Language": "",
        "LanguageISO": "",
        "AccountRelationships": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountRelationships"
            }
        },
        "AccountBalance": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountBalance"
            }
        },
        "AccountAddresses": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAddresses"
            }
        },
        "ContractAccounts": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/ContractAccounts"
            }
        },
        "PaymentCards": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/PaymentCards"
            }
        },
        "Correspondences": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/Correspondences"
            }
        },
        "CommunicationPreferences": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/CommunicationPreferences"
            }
        },
        "AccountContacts": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountContacts"
            }
        },
        "AccountType": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountType"
            }
        },
        "AccountAlerts": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAlerts"
            }
        },
        "AccountAddressIndependentEmails": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAddressIndependentEmails"
            }
        },
        "StandardAccountAddress": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/StandardAccountAddress"
            }
        },
        "AccountAddressIndependentMobilePhones": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAddressIndependentMobilePhones"
            }
        },
        "AccountAddressIndependentPhones": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAddressIndependentPhones"
            }
        },
        "AccountAddressIndependentFaxes": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountAddressIndependentFaxes"
            }
        },
        "PaymentDocuments": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/PaymentDocuments"
            }
        },
        "BankAccounts": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/BankAccounts"
            }
        },
        "AccountSex": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountSex"
            }
        },
        "AccountTitle": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/AccountTitle"
            }
        },
        "ServiceNotifications": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/ServiceNotifications"
            }
        },
        "ServiceOrders": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/ServiceOrders"
            }
        },
        "Outages": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/Outages"
            }
        },
        "Invoices": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/Invoices"
            }
        },
        "InteractionRecords": {
            "__deferred": {
                "uri": "http://s42020:8001/sap/opu/odata/sap/ERP_ISU_UMC/Accounts('2000018')/InteractionRecords"
            }
        }
    }
}`
	fmt.Fprintf(w, j)
}
func metadata(w http.ResponseWriter, r *http.Request){
	
	j := "metadata here we go"
	fmt.Fprintf(w, j)
}
var (
        username = "abc"
        password = "123"
)
func auth(w http.ResponseWriter, r *http.Request) {

        u, p, ok := r.BasicAuth()
        if !ok {
                fmt.Println("Error parsing basic auth")
                w.WriteHeader(401)
                return
        }
        if u != username {
                fmt.Printf("Username provided is correct: %s\n", u)
                w.WriteHeader(401)
                return
        }
        if p != password {
                fmt.Printf("Password provided is correct: %s\n", u)
                w.WriteHeader(401)
                return
        }
        fmt.Printf("Username: %s\n", u)
        fmt.Printf("Password: %s\n", p)
        w.WriteHeader(200)
        return
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/j", mockacct)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/sap/opu/odata/sap/ERP_ISU_UMC/$metadata",metadata)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
