module main

go 1.15

replace initialize => ../initialize

replace enc_dec => ../enc_dec

require (
	enc_dec v0.0.0-00010101000000-000000000000
	github.com/fentec-project/gofe v0.0.0-20210104123414-fd8f09f89d1c
	initialize v0.0.0-00010101000000-000000000000
)
