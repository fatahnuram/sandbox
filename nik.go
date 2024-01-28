package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getNikTahun(tahun int) (string, error) {
	tahunString := strconv.Itoa(tahun)
	tahunArr := strings.Split(tahunString, "")

	if len(tahunArr) != 4 {
		return "", errors.New("error: tahun harus 4 digit")
	}

	tahunLastDigitsArr := tahunArr[2:]
	tahunLastDigitsString := strings.Join(tahunLastDigitsArr, "")

	return tahunLastDigitsString, nil
}

func getNikGender(gender string) (string, error) {
	if gender != "ikhwan" && gender != "akhwat" {
		return "", errors.New("error: gender harus ikhwan atau akhwat")
	}

	mapGender := map[string]string{
		"ikhwan": "N",
		"akhwat": "T",
	}
	genderIdentifier := fmt.Sprint("AR", mapGender[gender])

	return genderIdentifier, nil
}

func getNikBulan(bulan int) (string, error) {
	if bulan < 1 || bulan > 12 {
		return "", errors.New("error: bulan harus antara 1 sampai 12")
	}

	nikBulan := "0"
	if bulan < 6 {
		nikBulan = "1"
	} else {
		nikBulan = "2"
	}

	return nikBulan, nil
}

func generateNIK(gender string, tahun int, bulan int, jumlah int) []string {
	nikGender, err := getNikGender(gender)
	if err != nil {
		log.Fatal(err)
	}

	nikTahun, err := getNikTahun(tahun)
	if err != nil {
		log.Fatal(err)
	}

	nikBulan, err := getNikBulan(bulan)
	if err != nil {
		log.Fatal(err)
	}

	nikPrefix := fmt.Sprint(nikGender, nikTahun, nikBulan)
	arrayNik := make([]string, jumlah)

	for i := 0; i < jumlah; i++ {
		arrayNik[i] = fmt.Sprintf("%v-%05d", nikPrefix, i+1)
	}
	fmt.Println(arrayNik)

	return arrayNik
}

func generateNIKLanjutan(lastNik string, jumlah int) ([]string, error) {
	if jumlah < 1 {
		return nil, errors.New("error: jumlah minimal 1")
	}

	nikOrigianlArr := strings.Split(lastNik, "-")
	if len(nikOrigianlArr) != 2 {
		return nil, errors.New("error: NIK malformed")
	}

	prefixNik := nikOrigianlArr[0]
	offsetNik := nikOrigianlArr[1]

	offset, err := strconv.Atoi(offsetNik)
	if err != nil {
		return nil, err
	}

	arrayNik := make([]string, jumlah)
	for i := 0; i < jumlah; i++ {
		arrayNik[i] = fmt.Sprintf("%v-%05d", prefixNik, offset+i+1)
	}
	fmt.Println(arrayNik)

	return arrayNik, nil
}
