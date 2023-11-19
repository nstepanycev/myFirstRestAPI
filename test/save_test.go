package test

import (
	"testing"
	"http"
	"fmt"
)

func TestSaveApi(t *testing.T){
	cases := struct{
		name string
		alias string
		url string
	}{
		name:  "Success",
        alias: "test_alias",
        url:   "https://google.com",
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T){
			urlSaverMock := mocks.NewURLSaver(t)
			if tc.respError == "" || tc.mockError != nil{
				urlSaverMock.On("SaveURL", tc.url, mock.AnythingOfType("string")).
				Return(int64(1), tc.mockError).
				Once()
			}
			handler := save.New(sl.NewDiscardLogger(), urlSaverMock)
			input := fmt.Sprintf(`{"url": "%s", "alias": "%s"}`, tc.url, tc.alias)
			req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader([]byte(input)))
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			require.Equal(t, rr.Code, http.StatusOK)
			body := rr.Body.String()

			var resp save.Response

			require.NoError(t, json.Unmarshal([]byte(body), &resp))
			require.Equal(t, tc.respError, resp.Error)

		})
	}
}