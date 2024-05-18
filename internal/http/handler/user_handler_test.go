package handler


func TestCreateUser(t *testing.T) {
    data := []byte(
    `{"first_name":"Kendrik","last_name":"Lamar","email":"kendrik.lamar@gmail.com","age":30}`
    )

    req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
}
