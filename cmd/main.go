package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"nexzipper/internal/archive"
)

func main() {
	fmt.Println("NexZipper'a hoş geldiniz!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Komutu girin (örn: compress, extract, help): ")
		command, _ := reader.ReadString('\n')
		command = trimNewline(command)
		command = strings.TrimSpace(command)

		switch command {
		case "nexzipper compress":
			fmt.Print("Kaynak dizinini girin: ")
			sourceDir, _ := reader.ReadString('\n')
			sourceDir = trimNewline(sourceDir)
			sourceDir = strings.TrimSpace(sourceDir)

			fmt.Print("Çıkış arşivi yolunu girin: ")
			outputFile, _ := reader.ReadString('\n')
			outputFile = trimNewline(outputFile)
			outputFile = strings.TrimSpace(outputFile)

			err := archive.Compress(sourceDir, outputFile)
			if err != nil {
				fmt.Println("Sıkıştırma hatası:", err)
			} else {
				fmt.Println("Sıkıştırma tamamlandı:", outputFile)
			}
		case "nexzipper extract":
			fmt.Print("Giriş arşivi yolunu girin: ")
			inputFile, _ := reader.ReadString('\n')
			inputFile = trimNewline(inputFile)
			inputFile = strings.TrimSpace(inputFile)

			fmt.Print("Çıkış dizinini girin: ")
			outputDir, _ := reader.ReadString('\n')
			outputDir = trimNewline(outputDir)
			outputDir = strings.TrimSpace(outputDir)

			err := archive.Extract(inputFile, outputDir)
			if err != nil {
				fmt.Println("Çıkartma hatası:", err)
			} else {
				fmt.Println("Çıkartma tamamlandı:", outputDir)
			}
		case "nexzipper help":
			printUsage()
		case "exit":
			fmt.Println("NexZipper kapatılıyor. İyi günler!")
			return
		default:
			fmt.Println("Geçersiz komut:", command)
		}
	}
}

func trimNewline(str string) string {
	return str[:len(str)-1]
}

func printUsage() {
	usage := `
Kullanım Kılavuzu:
compress   - Belirtilen kaynak dizini içeriğini çıkış arşivi dosyasına sıkıştırır.
extract    - Belirtilen giriş arşivi dosyasını çıkartarak çıkış dizinine çıkarır.
help       - Bu kılavuzu görüntüler.
exit       - Programı kapatır.
`
	fmt.Println(usage)
}
