# 🚀 Automation Smart Warehouse

**Automation Smart Warehouse** adalah platform berbasis Go untuk **monitoring infrastruktur warehouse**, yang dapat:
- Mendeteksi status layanan (sehat/tidak sehat)
- Menjalankan proses _self-healing_ secara otomatis
- Menyediakan **dashboard web interaktif** untuk melihat status audit real-time

---

## 📁 Struktur Proyek

```
automation-smart-warehouse/
│
├── cmd/
│ └── server/ # Entry point aplikasi
│ └── main.go
│
├── internal/
│ ├── config/ # Loader konfigurasi (.env)
│ ├── engine/ # Scheduler & Executor untuk audit/healing
│ ├── audit/ # Simulasi audit layanan
│ ├── service/ # Logika penyembuhan (self-healing)
│ ├── web/ # HTTP handler dan static HTML dashboard
│ └── static/
│ └── index.html # UI Dashboard Monitoring
│
├── rules/
│ └── rules.yaml # Placeholder untuk aturan audit
│
├── .env # Konfigurasi env
── go.mod / go.sum # Dependency
```


## ⚙️ Fitur Utama
```

| Fitur                      | Penjelasan                                                                 |
|---------------------------|---------------------------------------------------------------------------|
| 🔍 Audit Services          | Memindai status layanan (contoh: database, API) setiap 10 detik           |
| 🛠️ Self-Healing           | Jika layanan bermasalah, sistem akan otomatis menjalankan pemulihan       |
| 📊 Dashboard Web          | Tampilan visual status layanan secara real-time via HTML/JS               |
| 🔧 Rules via YAML         | (Optional) Gunakan `rules.yaml` untuk definisi aturan custom              |

```

## 🧪 Simulasi Audit

Output simulasi audit di terminal:
```
Scanning services...
Service api is unhealthy. Triggering remediation...
Restarting API service...
```

Dashboard Web akan menampilkan status:
- ✅ **Healthy**
- ❌ **Unhealthy**

---

## 🚀 Menjalankan Proyek

### 1. Clone & Setup

```bash
git clone https://github.com/ffahriza/automation-smart-warehouse.git
cd automation-smart-warehouse
```

### 2. Env.
```
MONGO_URI=mongodb://localhost:27017
DB_NAME=warehouse
RULES_PATH=./rules/rules.yaml
```

🧠 Teknologi yang Digunakan
```
Golang (Go 1.21+)
HTML5 + Vanilla JS
Environment Config (github.com/joho/godotenv)
Modular folder layout dengan internal/ pattern
```

📌 Catatan Developer
```
Endpoint API: GET /api/audit
Logika audit dan self-healing bisa dimodifikasi via internal/audit/ dan internal/service/
Untuk animasi status di dashboard, edit script.js dan style.css
```

🧑‍💻 Author
Ffahriza
GitHub: @ffahriza
