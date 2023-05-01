package crypto

type CryptoService interface {
	CreatePasswordHash(plainPassword string) (hashedPassword string, err error)
	ValidatePassword(hashedPassword, plainPassword string) (isValid bool)
	CreateMD5Hash(plainText string) (hashedText string)
}
