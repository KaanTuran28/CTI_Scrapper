# CTI Web Scraper

![Go](https://img.shields.io/badge/Go-1.25-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

<p align="center"><b><a href="#english">English</a></b> · <b><a href="#türkçe">Türkçe</a></b></p>

---

## English

A command-line web scraper written in Go for collecting data from websites in Cyber Threat Intelligence (CTI) workflows.

### What it does

It drives a real headless browser with the [chromedp](https://github.com/chromedp/chromedp) library, so it fully loads JavaScript-heavy, dynamic sites (Twitter, Reddit, etc.). For a target URL it takes a screenshot, saves the HTML source and extracts and reports every link on the page.

### Requirements

- Go 1.25+
- Chromium / Google Chrome (used by chromedp)

### Usage

```bash
go run main.go -url=https://target-site.com
go run main.go -url=https://target-site.com -headless=false   # show the browser window
```

> Use it only against sites you are authorised to scrape, and respect each site's terms of service and robots rules.

### License

MIT — see [LICENSE](./LICENSE).

---

## Türkçe

Siber Tehdit İstihbaratı (CTI) süreçlerinde web sitelerinden otomatik veri toplamak için Go diliyle yazılmış bir komut satırı aracı.

### Ne yapar?

[chromedp](https://github.com/chromedp/chromedp) kütüphanesiyle gerçek bir headless tarayıcı sürer; böylece JavaScript içeren dinamik siteleri (Twitter, Reddit vb.) eksiksiz yükler. Hedef URL için ekran görüntüsü alır, HTML kaynak kodunu kaydeder ve sayfadaki tüm linkleri ayıklayıp raporlar.

### Gereksinimler

- Go 1.25+
- Chromium / Google Chrome (chromedp kullanır)

### Kullanım

```bash
go run main.go -url=https://hedef-site.com
go run main.go -url=https://hedef-site.com -headless=false   # tarayıcı penceresini göster
```

> Yalnızca veri toplamaya yetkili olduğunuz sitelerde kullanın ve her sitenin kullanım şartlarına ve robots kurallarına uyun.

### Lisans

MIT — bkz. [LICENSE](./LICENSE).
