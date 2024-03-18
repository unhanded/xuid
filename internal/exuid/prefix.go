package exuid

type ExUIDAsciiPrefix struct {
	// A is the first byte of the ExUID, displayed as the corresponding ASCII character
	A byte
	// B is the second byte of the ExUID, displayed as the corresponding ASCII character
	B byte
	// C is the third byte of the ExUID, displayed as the corresponding ASCII character
	C byte
	// D is the fourth byte of the ExUID, displayed as the corresponding ASCII character
	// This is the last byte of the prefix
	D byte
}
