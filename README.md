# CTI Web Scraper

Bu proje, Siber Tehdit İstihbaratı (CTI) süreçlerinde web sitelerinden otomatik veri toplamak amacıyla Go diliyle geliştirilmiştir.

**Ne Yapar?**
Program, Chromedp kütüphanesini kullanarak gerçek bir tarayıcı simülasyonu yapar. Bu sayede JavaScript içeren dinamik siteleri (Twitter, Reddit vb.) eksiksiz yükler; anlık ekran görüntüsünü alır, HTML kaynak kodunu kaydeder ve site içindeki tüm linkleri ayıklayarak raporlar.

**Nasıl Kullanılır?**
Projeyi indirdikten sonra terminalde şu komutu kullanmanız yeterlidir:

`go run main.go -url=https://www.hedef-site.com`
