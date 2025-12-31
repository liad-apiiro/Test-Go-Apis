package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Standard net/http server with various methods and non-standard methods
func setupNetHTTPServer() {
	// Standard GET endpoint
	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"users": [{"id": 1, "name": "Alice"}]}`))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Standard POST endpoint
	http.HandleFunc("/api/v1/users/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := io.ReadAll(r.Body)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"message": "User created", "data": %s}`, string(body))))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Multiple methods on same endpoint
	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"products": [{"id": 1, "name": "Widget"}]}`))
		case http.MethodPost:
			body, _ := io.ReadAll(r.Body)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"message": "Product created", "data": %s}`, string(body))))
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Product updated"}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Product deleted"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Non-standard HTTP method: FUNKYTOWN
	http.HandleFunc("/api/v1/funkytown", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "FUNKYTOWN" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Welcome to Funkytown!"}`))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Another non-standard method: DANCE
	http.HandleFunc("/api/v1/dance", func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) == "DANCE" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Let's dance!"}`))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Multiple non-standard methods on same endpoint
	http.HandleFunc("/api/v1/custom", func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToUpper(r.Method)
		switch method {
		case "FUNKYTOWN":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "funkytown", "status": "groovy"}`))
		case "DANCE":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "dance", "status": "moving"}`))
		case "PARTY":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "party", "status": "celebrating"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Discouraged: Missing method check - accepts any method
	http.HandleFunc("/api/v1/bad/no-method-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "No method validation - bad practice!"}`))
	})

	// Discouraged: Missing error handling
	http.HandleFunc("/api/v1/bad/no-error-handling", func(w http.ResponseWriter, r *http.Request) {
		// No error handling for body read
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// Discouraged: Missing Content-Type header
	http.HandleFunc("/api/v1/bad/no-content-type", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Missing Content-Type header
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"data": "missing content-type"}`))
		}
	})

	// Discouraged: Insecure - no input validation
	http.HandleFunc("/api/v1/bad/no-validation", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := io.ReadAll(r.Body)
			// No validation of input
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"echo": %s}`, string(body))))
		}
	})

	fmt.Println("Net/HTTP server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func TestTakenFromRealCase(t *testing.T) {
	// timeNow = mockTimeNow

	mux := http.NewServeMux()

	loginHandler := newHandler(func(w http.ResponseWriter, r *http.Request, _ int32) {
		// r.ParseForm()

		// username := r.Form.Get("j_username")
		// password := r.Form.Get("j_password")

		// require.Equal(t, "username", username)
		// require.Equal(t, "password", password)
		w.WriteHeader(http.StatusOK)
	})

	tokenHandler := newHandler(func(w http.ResponseWriter, _ *http.Request, _ int32) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("testtoken"))
	})

	endpointHandler := newHandler(func(w http.ResponseWriter, r *http.Request, _ int32) {
		token := r.Header.Get("X-XSRF-TOKEN")

		// Ensure token is correctly passed as header
		// require.Equal(t, "testtoken", token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	})

	mux.HandleFunc("/j_security_check", loginHandler.Func)
	mux.HandleFunc("/dataservice/client/token", tokenHandler.Func)
	mux.HandleFunc("/dataservice/device", endpointHandler.Func)
	mux.HandleFunc(fmt.Sprintf("%s /Blip/Bloop", http.MethodGet), endpointHandler.Func)
	mux.HandleFunc("POST /Flip/Flap", endpointHandler.Func)

	server := httptest.NewServer(mux)
	defer server.Close()

	// client, err := NewClient(serverURL(server), "username", "password", true)
	// require.NoError(t, err)

	// _, err = client.GetDevices()
	// require.NoError(t, err)

	// require.Equal(t, "testtoken", client.token, "Token should be set correctly")
	// require.Equal(t, int64(946688400), client.tokenExpiry.Unix(), "Token expiry should be set correctly")

	// // Ensure login endpoint has been called 1 times
	// require.Equal(t, 1, loginHandler.numberOfCalls())
	// // Ensure token endpoint has been called 1 times
	// require.Equal(t, 1, loginHandler.numberOfCalls())

	// Re-call GetDevices and ensure auth is not re-called
	// _, err = client.GetDevices()
	// require.NoError(t, err)

	// // Ensure login endpoint has been called 1 times
	// require.Equal(t, 1, loginHandler.numberOfCalls())
	// // Ensure token endpoint has been called 1 times
	// require.Equal(t, 1, loginHandler.numberOfCalls())

	// Fast-forward 1h01 and ensure token is correctly expired
	// timeNow = func() time.Time {
	// 	return mockTimeNow().Add(time.Hour + time.Minute)
	// }

	// Re-call GetDevices and ensure auth is re-called
	// _, err = client.GetDevices()
	// require.NoError(t, err)

	// // Ensure login endpoint has been called 2 times
	// require.Equal(t, 2, loginHandler.numberOfCalls())
	// // Ensure token endpoint has been called 2 times
	// require.Equal(t, 2, loginHandler.numberOfCalls())
}