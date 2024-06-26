package cli

import (
	"fmt"
	"zorro/db"
)

var zorro = "\u001b[32m" +
	`
__________
\____    /__________________  ____                        /\                               ______,....----,
  /     //  _ \_  __ \_  __ \/  _ \         /VVVVVVVVVVVVVV|===================""""""""""""       ___,..-' 
 /     /(  <_> )  | \/|  | \(  <_> )        \^^^^^^^^^^^^^^|======================----------""""""
/_______ \____/|__|   |__|   \____/                       \/
\/             
` +
	"\u001b[0m"

// "\u001b[34m" blue
var gopher = "\u001B[32m" +
	`
         ,_---~~~~~----._         
  _,,_,*^____      _____*g*\"*,@\
 / __/ /'     ^.  /      \ ^@q   ] 
[  @f | @))    |  | @))   l  0 _/  
 \ /   \~____ / __ \_____/    \
|           _l__l_           I
}          [______]           I
]            | | |            |
]             ~ ~             |
|                            |
` + "\n" +
	"\u001b[0m"

func Help() {
	fmt.Println(zorro, gopher)
	fmt.Println("\nUsage:")
	fmt.Println("  zorro <command> [args]")
	fmt.Println("\nCommands:\n")
	fmt.Println("  encrypt <input_file> <output_file> - Encrypt a file with AES")
	fmt.Println("  decrypt <input_file> <output_file> - Decrypt a file with AES")
	fmt.Println("  rsa-key generate <bit_length> [<pem_file>] [--save <key_name>] - Generate RSA key")
	fmt.Println("  rsa-key list - List RSA keys in the database ")
	fmt.Println("  hash <method> <value> - Compute hash of a value using specified method")
	fmt.Println("\nHash Methods:")
	fmt.Println("  - sha256: Compute SHA-256 hash")
	fmt.Println("  - md5: Compute MD5 hash")
}

func processListDb() {
	fmt.Println(zorro)
	db.GetRSAKeys()
}

func helpHash() {
	fmt.Println("Usage: zorro hash <method> <value>")
	fmt.Println("  Compute hash of a value using specified method")
	fmt.Println("\nAvailable Hash Methods:")
	fmt.Println("  - sha256: Compute SHA-256 hash")
	fmt.Println("  - md5: Compute MD5 hash")
}

func helpEncrypt() {
	fmt.Print(zorro)
	fmt.Println("\nUsage: zorro encrypt <input> <output_file>")
	fmt.Println("  Encrypts a file or string.")
	fmt.Println("  <input> can be a file path or a string.")
	fmt.Println("\nExample:")
	fmt.Println("  zorro encrypt file.txt encrypted.txt  # Encrypts the content of the file 'file.txt'")
	fmt.Println("  zorro encrypt 'hello' encrypted.txt   # Encrypts the string 'hello'")
}

func helpDecrypt() {
	fmt.Print(zorro)
	fmt.Println("\nUsage: zorro decrypt <input> <output_file>")
	fmt.Println("  Decrypts a file or string.")
	fmt.Println("  <input> can be a file path or a string.")
	fmt.Println("\nExample:")
	fmt.Println("  zorro decrypt encrypted.txt decrypted.txt  # Decrypts the content of the file 'encrypted.txt'")
	fmt.Println("  zorro decrypt 's3cr3t' decrypted.txt        # Decrypts the string 's3cr3t'")
}

func helpRSAKeyGenerate() {
	fmt.Print(zorro)
	fmt.Println("\nUsage: zorro rsa-key generate <bit_length> [<pem_file>] [--save <key_name>]")
	fmt.Println("  Generates an RSA key pair.")
	fmt.Println("  Options:")
	fmt.Println("    <bit_length>          : The length of the key in bits (e.g., 2048, 4096)")
	fmt.Println("    [<pem_file>]          : Optional. The name of the PEM file to save the key.")
	fmt.Println("                            If not provided, the default name 'private_key.pem' will be used.")
	fmt.Println("    [--save <key_name>]: Optional. Save the key to the database with the specified name.")
	fmt.Println("\nExample:")
	fmt.Println("  zorro rsa-key generate 2048 private.pem        # Generates a 2048-bit RSA key pair and saves to 'private.pem'")
	fmt.Println("  zorro rsa-key generate 4096                    # Generates a 4096-bit RSA key pair with default name: private_key.pem")
}
