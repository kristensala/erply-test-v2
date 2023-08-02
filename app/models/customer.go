package models

type CreateCustomerRequest struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Code string `json:"code" binding:"required"`
    Email string `json:"email"`
    CompanyName string `json:"companyName" binding:"required"`

}

type UpdateCustomerRequest struct {
    Id string `json:"id: binding:"required"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Code string `json:"code"`
    Email string `json:"email"`
}

type CustomerResponse struct {
    ID                      int                         `json:"id"`
    PayerID                 int                         `json:"payerID,omitempty"`
    CustomerID              int                         `json:"customerID"`
    TypeID                  string                      `json:"type_id"`
    FullName                string                      `json:"fullName"`
    CompanyName             string                      `json:"companyName"`
    FirstName               string                      `json:"firstName"`
    LastName                string                      `json:"lastName"`
    GroupID                 int                         `json:"groupID"`
    EDI                     string                      `json:"EDI"`
    GLN                     string                      `json:"GLN"`
    IsPOSDefaultCustomer    int                         `json:"isPOSDefaultCustomer"`
    CountryID               string                      `json:"countryID"`
    Phone                   string                      `json:"phone"`
    EInvoiceEmail           string                      `json:"eInvoiceEmail"`
    Email                   string                      `json:"email"`
    Fax                     string                      `json:"fax"`
    Code                    string                      `json:"code"`
    ReferenceNumber         string                      `json:"referenceNumber"`
    VatNumber               string                      `json:"vatNumber"`
    BankName                string                      `json:"bankName"`
    BankAccountNumber       string                      `json:"bankAccountNumber"`
    BankIBAN                string                      `json:"bankIBAN"`
    BankSWIFT               string                      `json:"bankSWIFT"`
    PaymentDays             int                         `json:"paymentDays"`
    Notes                   string                      `json:"notes"`
    LastModified            int                         `json:"lastModified"`
    CustomerType            string                      `json:"customerType"`
    Address                 string                      `json:"address"`
    Street                  string                      `json:"street"`
    Address2                string                      `json:"address2"`
    City                    string                      `json:"city"`
    PostalCode              string                      `json:"postalCode"`
    Country                 string                      `json:"country"`
    State                   string                      `json:"state"`
    Credit                  int                         `json:"credit"`
    CompanyTypeID           int                         `json:"companyTypeID"`
    PersonTitleID           int                         `json:"personTitleID"`
    EmailEnabled            int                         `json:"emailEnabled"`
    MailEnabled             int                         `json:"mailEnabled"`
    EInvoiceEnabled         int                         `json:"eInvoiceEnabled"`
    FlagStatus              int                         `json:"flagStatus"`
    OperatorIdentifier      string                      `json:"operatorIdentifier"`
    Gender                  string                      `json:"gender"`
    GroupName               string                      `json:"groupName"`
    Mobile                  string                      `json:"mobile"`
    Birthday                string                      `json:"birthday"`
    IntegrationCode         string                      `json:"integrationCode"`
    ColorStatus             string                      `json:"colorStatus"`
    FactoringContractNumber string                      `json:"factoringContractNumber"`
    Image                   string                      `json:"image"`
    TwitterID               string                      `json:"twitterID"`
    FacebookName            string                      `json:"facebookName"`
    CreditCardLastNumbers   string                      `json:"creditCardLastNumbers"`
    EuCustomerType          string                      `json:"euCustomerType"`
    CustomerCardNumber      string                      `json:"customerCardNumber"`
    LastModifierUsername    string                      `json:"lastModifierUsername"`
    DefaultAssociationName  string                      `json:"defaultAssociationName"`
    DefaultProfessionalName string                      `json:"defaultProfessionalName"`
    TaxExempt               int                         `json:"taxExempt"`
    PaysViaFactoring        int                         `json:"paysViaFactoring"`
    SalesBlocked            int                         `json:"salesBlocked"`
    RewardPointsDisabled    int                         `json:"rewardPointsDisabled"`
    CustomerBalanceDisabled int                         `json:"customerBalanceDisabled"`
    PosCouponsDisabled      int                         `json:"posCouponsDisabled"`
    EmailOptOut             int                         `json:"emailOptOut"`
    ShipGoodsWithWaybills   int                         `json:"shipGoodsWithWaybills"`
    DefaultAssociationID    int                         `json:"defaultAssociationID"`
    DefaultProfessionalID   int                         `json:"defaultProfessionalID"`

    // Web-shop related fields
    Username  string `json:"webshopUsername"`
    LastLogin string `json:"webshopLastLogin"`

    // Detailed info
    PriceListID  int `json:"priceListID"`
    PriceListID2 int `json:"priceListID2"`
    PriceListID3 int `json:"priceListID3"`
}
