package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SilentControllers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}

func IndexPage(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query ke database
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Ambil kolom nama
		columns, err := rows.Columns()
		if err != nil {
			http.Error(w, "Error retrieving columns", http.StatusInternalServerError)
			return
		}

		// Simpan hasil query
		var results []map[string]interface{}
		for rows.Next() {
			// Buat slice untuk menyimpan nilai kolom
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range columns {
				valuePtrs[i] = &values[i]
			}

			// Scan data ke dalam nilai
			if err := rows.Scan(valuePtrs...); err != nil {
				http.Error(w, "Error scanning row", http.StatusInternalServerError)
				return
			}

			// Buat map untuk menyimpan data baris
			rowData := make(map[string]interface{})
			for i, col := range columns {
				val := values[i]
				if b, ok := val.([]byte); ok {
					rowData[col] = string(b)
				} else {
					rowData[col] = val
				}
			}

			results = append(results, rowData)
		}

		// Buat respons JSON
		response := JsonResponse{
			Status:  true,
			Message: "Data retrieved successfully",
			Data:    results,
		}

		// Set Content-Type ke application/json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode dan kirim respons JSON
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}
}
