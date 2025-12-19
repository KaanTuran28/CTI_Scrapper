package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/yosssi/gohtml"
)

func main() {
	siteAdresi := flag.String("url", "", "Hedef site adresi")
	flag.Parse()

	if *siteAdresi == "" {
		fmt.Println("Hata: Lütfen bir URL girin. Örnek: -url=https://www.google.com")
		return
	}

	u, err := url.Parse(*siteAdresi)
	if err != nil {
		log.Fatal("URL hatalı:", err)
	}

	domain := strings.TrimPrefix(u.Hostname(), "www.")
	if domain == "" {
		domain = "site_verisi"
	}

	fmt.Println("Bağlanılıyor:", *siteAdresi)

	ayarlar := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.WindowSize(1920, 1080),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), ayarlar...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlIcerik string
	var ekranGoruntusu []byte

	fmt.Println("Veriler çekiliyor...")
	err = chromedp.Run(ctx,
		chromedp.Navigate(*siteAdresi),
		chromedp.Sleep(3*time.Second), 
		chromedp.CaptureScreenshot(&ekranGoruntusu),
		chromedp.OuterHTML("html", &htmlIcerik),
	)

	if err != nil {
		log.Fatal("İşlem sırasında hata oluştu:", err)
	}

	resimDosyasi := domain + "_screenshot.png"
	if err := os.WriteFile(resimDosyasi, ekranGoruntusu, 0644); err != nil {
		log.Fatal("Resim kaydedilemedi:", err)
	}
	fmt.Println("- Ekran görüntüsü kaydedildi:", resimDosyasi)

	txtDosyasi := domain + "_data.txt"
	duzenliHtml := gohtml.Format(htmlIcerik)

	if err := os.WriteFile(txtDosyasi, []byte(duzenliHtml), 0644); err != nil {
		log.Fatal("Veri dosyası kaydedilemedi:", err)
	}
	fmt.Println("- HTML verisi kaydedildi:", txtDosyasi)

	linkDosyasi := domain + "_urls.txt"
	fmt.Println("- Linkler toplanıyor...")
	
	r := regexp.MustCompile(`href=["'](http[^"']+)["']`)
	bulunanLinkler := r.FindAllStringSubmatch(htmlIcerik, -1)

	if len(bulunanLinkler) > 0 {
		f, err := os.Create(linkDosyasi)
		if err != nil {
			log.Fatal("Link dosyası oluşturulamadı:", err)
		}
		defer f.Close()

		linkMap := make(map[string]bool)
		sayac := 0

		for _, m := range bulunanLinkler {
			urlLink := m[1]
			if !linkMap[urlLink] {
				linkMap[urlLink] = true
				f.WriteString(urlLink + "\n")
				sayac++
			}
		}
		fmt.Printf("- Toplam %d adet link bulundu ve '%s' dosyasına yazıldı.\n", sayac, linkDosyasi)
	} else {
		fmt.Println("- Sayfada hiç link bulunamadı.")
	}
	
	fmt.Println("İşlem tamamlandı.")
}