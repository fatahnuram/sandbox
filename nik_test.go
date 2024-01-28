package main

import "testing"

func TestGetNikTahunValid(t *testing.T) {
	input := 2019
	expected := "19"

	result, err := getNikTahun(input)

	if result != expected || err != nil {
		t.Fatalf("getNikTahun, result=%v, expected=%v, err=%v", result, expected, err)
	}
}

func TestGetNikTahunLessThan4Digit(t *testing.T) {
	input := 201

	result, err := getNikTahun(input)

	if err == nil {
		t.Fatalf("getNikTahun, result=%v, err=%v", result, err)
	}
}

func TestGetNikTahunMoreThan4Digit(t *testing.T) {
	input := 20142

	result, err := getNikTahun(input)

	if err == nil {
		t.Fatalf("getNikTahun, result=%v, err=%v", result, err)
	}
}

func TestGetNikGenderAkhwatValid(t *testing.T) {
	input := "akhwat"
	expected := "ART"

	result, err := getNikGender(input)

	if err != nil || result != expected {
		t.Fatalf("getNikGender, result=%v, expected=%v, err=%v", result, expected, err)
	}
}

func TestGetNikGenderIkhwanValid(t *testing.T) {
	input := "ikhwan"
	expected := "ARN"

	result, err := getNikGender(input)

	if err != nil || result != expected {
		t.Fatalf("getNikGender, result=%v, expected=%v, err=%v", result, expected, err)
	}
}

func TestGetNikGenderInvalid(t *testing.T) {
	input := "perempuan"

	result, err := getNikGender(input)

	if err == nil {
		t.Fatalf("getNikGender, result=%v, err=%v", result, err)
	}
}

func TestGetNikBulanSemesterSatu(t *testing.T) {
	input := 3
	expected := "1"

	result, err := getNikBulan(input)

	if err != nil || result != expected {
		t.Fatalf("getNikBulan, result=%v, expected=%v, err=%v", result, expected, err)
	}
}

func TestGetNikBulanSemesterDua(t *testing.T) {
	input := 6
	expected := "2"

	result, err := getNikBulan(input)

	if err != nil || result != expected {
		t.Fatalf("getNikBulan, result=%v, expected=%v, err=%v", result, expected, err)
	}
}

func TestGetNikBulanMoreThan12(t *testing.T) {
	input := 13

	result, err := getNikBulan(input)

	if err == nil {
		t.Fatalf("getNikBulan, result=%v, err=%v", result, err)
	}
}

func TestGetNikBulanLessThan1(t *testing.T) {
	input := 0

	result, err := getNikBulan(input)

	if err == nil {
		t.Fatalf("getNikBulan, result=%v, err=%v", result, err)
	}
}

func TestGenerateNIKAkhwat(t *testing.T) {
	expected := []string{
		"ART192-00001",
		"ART192-00002",
		"ART192-00003",
	}

	result := generateNIK("akhwat", 2019, 9, 3)

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("generateNIK, result=%v, expected=%v", result, expected)
		}
	}
}

func TestGenerateNIKIkhwan(t *testing.T) {
	expected := []string{
		"ARN241-00001",
		"ARN241-00002",
		"ARN241-00003",
	}

	result := generateNIK("ikhwan", 2024, 2, 3)

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("generateNIK, result=%v, expected=%v", result, expected)
		}
	}
}

func TestGenerateNIKLanjutanValid(t *testing.T) {
	expected := []string{
		"ARN201-00036",
		"ARN201-00037",
		"ARN201-00038",
	}

	result, err := generateNIKLanjutan("ARN201-00035", 3)

	if err != nil {
		t.Fatalf("generateNIKLanjutan, result=%v, expected=%v, err=%v", result, expected, err)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("generateNIKLanjutan, result=%v, expected=%v", result, expected)
		}
	}
}

func TestGenerateNIKLanjutanNIKInvalidNoSeparator(t *testing.T) {
	result, err := generateNIKLanjutan("ARN20100035", 3)

	if err == nil {
		t.Fatalf("generateNIKLanjutan, result=%v, err=%v", result, err)
	}
}

func TestGenerateNIKLanjutanNIKInvalidNonInteger(t *testing.T) {
	result, err := generateNIKLanjutan("ARN201-abcde", 3)

	if err == nil {
		t.Fatalf("generateNIKLanjutan, result=%v, err=%v", result, err)
	}
}

func TestGenerateNIKLanjutanGenerateZero(t *testing.T) {
	result, err := generateNIKLanjutan("ARN201-00035", 0)

	if err == nil {
		t.Fatalf("generateNIKLanjutan, result=%v, err=%v", result, err)
	}
}

func TestGenerateNIKLanjutanGenerateMinus(t *testing.T) {
	result, err := generateNIKLanjutan("ARN201-00035", -5)

	if err == nil {
		t.Fatalf("generateNIKLanjutan, result=%v, err=%v", result, err)
	}
}
