package api

import (

)
// Ошибки перенести
// func Error(msg string) Response{
// 	return Response{
// 		Status: StatusError,
// 		Error: msg,
// 	}
// }

// func OK() Response{
// 	return Response{
// 		Status: StatusOK,
// 	}
// }

// func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
//     render.JSON(w, r, Response{
//         Response: resp.OK(),
//         Alias:    alias,
//     })
// }