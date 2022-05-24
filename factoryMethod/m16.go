package main

type m16 struct {
	gun
}

func newM16() iGun {
	return &m16{
		gun{
			name:  "m16",
			power: 1,
		},
	}
}
