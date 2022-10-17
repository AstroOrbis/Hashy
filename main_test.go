package main

import "testing"

type testCase struct {
	input          string
	expectedOutput string
}

func TestMD5hash(t *testing.T) {
	MD5TestCases := []testCase{
		{
			input:          "",
			expectedOutput: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			input:          "anfljavkjaswiGJAIEORPNGAbfgjasfbvkja valjnnfgwiegpqo.  snfaonJIGOSNFisonfgIIfbnnifgPInPINnpiwprgqbwrpug",
			expectedOutput: "cf0f4dadbf16adef16071015b79bbe56",
		},
		{
			input: ` fvjkfvfhvp aifwirhfwpeh 4y8925047581y0 gfbvsbgu0yt74y7834giuf g


hfoufgh y02475y gfugapfgha fg`,
			expectedOutput: "7657caa293758ce9141058d2d5a0d465",
		},
	}

	for _, tc := range MD5TestCases {
		if op := MD5hash(tc.input); op != tc.expectedOutput {
			t.Errorf("ERROR: MD5hash returned %v, expected %v", op, tc.expectedOutput)
		}
	}
}

func TestSHA1hash(t *testing.T) {
	SHA1TestCases := []testCase{
		{
			input:          "",
			expectedOutput: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			input:          "anfljavkjaswiGJAIEORPNGAbfgjasfbvkja valjnnfgwiegpqo.  snfaonJIGOSNFisonfgIIfbnnifgPInPINnpiwprgqbwrpug",
			expectedOutput: "93359d8f59042d6899b669917d86183acd9bf59b",
		},
		{
			input: ` fvjkfvfhvp aifwirhfwpeh 4y8925047581y0 gfbvsbgu0yt74y7834giuf g


hfoufgh y02475y gfugapfgha fg`,
			expectedOutput: "e6f138b899fedf6770c573c704617cb1fabf0fc9",
		},
	}
	for _, tc := range SHA1TestCases {
		if op := SHA1hash(tc.input); op != tc.expectedOutput {
			t.Errorf("ERROR: SHA1hash returned %v, expected %v", op, tc.expectedOutput)
		}
	}
}

func TestSHA256hash(t *testing.T) {
	SHA256TestCases := []testCase{
		{
			input:          "",
			expectedOutput: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			input:          "anfljavkjaswiGJAIEORPNGAbfgjasfbvkja valjnnfgwiegpqo.  snfaonJIGOSNFisonfgIIfbnnifgPInPINnpiwprgqbwrpug",
			expectedOutput: "58da92cbc2fcadc1ebbfab7f8f85c32211b410602c53eb4eacee9b2804d09bf8",
		},
		{
			input: ` fvjkfvfhvp aifwirhfwpeh 4y8925047581y0 gfbvsbgu0yt74y7834giuf g


hfoufgh y02475y gfugapfgha fg`,
			expectedOutput: "c78c03d6423aa531ce2d12519c9a9cc2fb367e5aec4c8f213b008c6efc2d39ee",
		},
	}
	for _, tc := range SHA256TestCases {
		if op := SHA256hash(tc.input); op != tc.expectedOutput {
			t.Errorf("ERROR: SHA256hash returned %v, expected %v", op, tc.expectedOutput)
		}
	}
}

func TestSHA512hash(t *testing.T) {
	SHA512TestCases := []testCase{
		{
			input:          "",
			expectedOutput: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
		{
			input:          "anfljavkjaswiGJAIEORPNGAbfgjasfbvkja valjnnfgwiegpqo.  snfaonJIGOSNFisonfgIIfbnnifgPInPINnpiwprgqbwrpug",
			expectedOutput: "3ce3c3a6c4a96b5eb1288eefc445e801cced23c84c86cad7df6805d567d8c304a9143eb16d670d5dbed2037eb9166d598797108bb81030b1187027d747f1ad9a",
		},
		{
			input: ` fvjkfvfhvp aifwirhfwpeh 4y8925047581y0 gfbvsbgu0yt74y7834giuf g


hfoufgh y02475y gfugapfgha fg`,
			expectedOutput: "3bd813da98d3a3ad131370c3498f450538ad82a83054a9716104015fb0a9e765125430d6800c76a4054c042b4596975b3d25e6fed165ba0af92c417c339b4e28",
		},
	}
	for _, tc := range SHA512TestCases {
		if op := SHA512hash(tc.input); op != tc.expectedOutput {
			t.Errorf("ERROR: SHA512hash returned %v, expected %v", op, tc.expectedOutput)
		}
	}
}
