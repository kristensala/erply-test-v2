package models

type CreateCustomerRequest struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Code string `json:"code" binding:"required"`
    Email string `json:"email"`
    CompanyName string `json:"companyName" binding:"required"`

}

type UpdateCustomerRequest struct {
    Id string `json:"id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Code string `json:"code"`
    Email string `json:"email"`
}

type CustomerResponse struct {
    Id int `json:"id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Code string `json:"code"`
    Email string `json:"email"`
    CompanyName string `json:"companyName"`
    Address string `json:"address"`
}
