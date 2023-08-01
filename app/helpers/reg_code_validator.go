package helpers

import "strconv"

func IsValidRegCode(code string) bool {
    if len(code) != 11 {
        return false
    }

    lastNr, _ := strconv.Atoi(string(code[len(code) - 1]))
    controlNumber := checkControlNumber(code)

    if lastNr == controlNumber {
        return true
    }

    return false
}

func checkControlNumber(code string) int {
    multypliersPrimary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
    multypliersSecondary := []int{3, 4, 5, 6, 7, 8, 9, 1, 2, 3}

    var sum int
    for index, multyplier := range multypliersPrimary {
        number, _ := strconv.Atoi(string(code[index]))
        sum = sum + (multyplier * number)
    }

    controlNumberPrimary := sum % 11

    if controlNumberPrimary < 10 {
        return controlNumberPrimary
    }

    if controlNumberPrimary == 10 {
        var multypliersSecondarySum int
        for index, multyplier := range multypliersSecondary {
            number, _ := strconv.Atoi(string(code[index]))
            multypliersSecondarySum = multypliersSecondarySum + (multyplier * number)
        }

        controlNumber := multypliersSecondarySum % 11
        if controlNumber < 10 {
            return controlNumber
        }

        if controlNumber == 10 {
            return 0
        }
    }

    return -1
}

