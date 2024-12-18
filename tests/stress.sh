#!/bin/bash

# URL endpoint
URL="http://localhost:8000/users"

# Loop untuk mengirim 100 permintaan
for i in {1..10000}; do
    # Data JSON yang berubah-ubah
    JSON=$(cat <<EOF
{
    "email": "user$i@example.com",
    "password": "secretpassword"
}
EOF
    )
    # Kirim permintaan menggunakan curl
    curl -X POST -H "Content-Type: application/json" -d "$JSON" "$URL"
    echo ""  # Tambahkan baris kosong untuk output yang lebih rapi
done
